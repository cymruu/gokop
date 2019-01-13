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
type HTMLError struct {
	Message string
	Code    int
}

func (h *HTMLError) Error() string {
	return fmt.Sprintf("HTMLCode (wtf?) Response Code: %d\nHTML:\n%s", h.Code, h.Message)
}

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
	request, err := req.ToHTTPRequest()
	if err != nil {
		return nil, err
	}
	res, err := w.httpClient.Do(request)
	if err != nil {

		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		err = &HTMLError{Message: string(data), Code: res.StatusCode}
		return nil, err
	}
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
	if len(postParams) > 0 {
		keys := make([]string, len(postParams))
		i := 0
		for k := range postParams {
			keys[i] = k
			i++
		}
		sort.Strings(keys)
		for _, v := range keys {
			tosign += postParams[v][0] + ","
		}
		tosign = tosign[:len(tosign)-1]
	}
	checksum := md5.Sum([]byte(tosign))
	req.GetHeaders().Add("apisign", fmt.Sprintf("%x", checksum))
}
