package gokop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		req.Sign(w)
	}
	request := req.ToHttpRequest()
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
