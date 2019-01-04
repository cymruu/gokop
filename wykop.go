package gokop

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

//DefaultUseragent is a default useragent used by wrapper clients
const DefaultUseragent = "gokop 0.0.1"

// //A function type to receive response
// type ResponseHandler func(errorTarget interface{}, target interface{})

//WykopAPI base struct
type WykopAPI struct {
	apikey string
	secret string

	httpClient *http.Client

	logger *log.Logger
}

func CreateAPIBase(apikey, secret string) WykopAPI {
	return WykopAPI{
		apikey:     apikey,
		secret:     secret,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		logger:     log.New(os.Stdout, ">>", 0),
	}
}
func (w *WykopAPI) APIKey() string {
	return w.apikey
}
func (w *WykopAPI) Secret() string {
	return w.secret
}
func (w *WykopAPI) SignRequest(req *WykopRequest) {
	tosign := w.secret + req.BuildURL()
	if req.postParams != nil {
		keys := make([]string, len(req.postParams))
		for k := range req.postParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, v := range keys {
			tosign += req.postParams[v][0] + ","
		}
		tosign = tosign[:len(tosign)-1]
	}
	w.logger.Println(tosign)
	checksum := md5.Sum([]byte(tosign))
	req.Header.Add("apisign", fmt.Sprintf("%x", checksum))
	w.logger.Println(req.Header.Get("apisign"))
}
func (w *WykopAPI) SendRequest(req *WykopRequest) ([]byte, error) {
	if !req.IsSigned() {
		w.SignRequest(req)
	}
	requestMethod := req.Method()
	request, _ := http.NewRequest(requestMethod.ToString(), req.BuildURL(), strings.NewReader(req.postParams.Encode()))
	request.Header = req.Header
	res, err := w.httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data, nil
}
func DecodeJSON(data []byte, target interface{}) error {
	switch v := target.(type) {
	case *string:
		*v = string(data)
		return nil
	default:
		return json.Unmarshal(data, target)
	}
}
