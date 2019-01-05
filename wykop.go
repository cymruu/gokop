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
	"time"
)

//DefaultUseragent is a default useragent used by wrapper clients
const DefaultUseragent = "gokop 0.0.1"

// //A function type to receive response
// type ResponseHandler func(errorTarget interface{}, target interface{})
type APIVersionT uint8

const (
	APIVersionV1 APIVersionT = 1
	APIVersionV2 APIVersionT = 2
)

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

func (w *WykopAPI) SendRequest(req IWykopRequest) ([]byte, error) {
	if !req.IsSigned() {
		w.Sign(req)
	}
	request, err := req.ToRequest()
	if err != nil {
		return nil, err
	}
	res, err := w.httpClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	return data, err
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
func (w *WykopAPI) Sign(req IWykopRequest) {
	tosign := w.Secret() + req.BuildURL()
	postParams := req.GetPostParams()
	if postParams != nil {
		keys := make([]string, len(postParams))
		for k := range postParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, v := range keys {
			tosign += postParams[v][0] + ","
		}
		tosign = tosign[:len(tosign)-1]
	}
	fmt.Println(tosign)
	checksum := md5.Sum([]byte(tosign))
	req.GetHeaders().Add("apisign", fmt.Sprintf("%x", checksum))
	fmt.Println(req.GetHeaders().Get("apisign"))
}
