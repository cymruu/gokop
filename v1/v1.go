package v1

import (
	"fmt"

	"github.com/cymruu/gokop"
	"github.com/cymruu/gokop/v1/models"
)

const (
	_userAgent = "gokop 0.0.1 v1"
	apiVersion = "v1"
	APIURL     = "https://a.wykop.pl"
	//response types
	responseJSON  = "json"
	responseJSONP = "jsonp"
	responseXML   = "xml"

	//response filters
	outputCLEAR = "clear"
)

type WykopAPIV1 struct {
	*gokop.WykopAPI

	userkey string

	useragent string
	baseURL   string
}

func (w *WykopAPIV1) Useragent() string {
	userAgent := gokop.DefaultUseragent
	if w.userkey != "" {
		userAgent = w.useragent
	}
	return userAgent
}
func (w *WykopAPIV1) SetUseragent(useragent string) {
	w.useragent = useragent
}
func (w *WykopAPIV1) APIVersion() string {
	return apiVersion
}
func (w *WykopAPIV1) Userkey() string {
	return w.userkey
}
func (w *WykopAPIV1) SetUserkey(userkey string) {
	w.userkey = userkey
}

func CreateWykopV1API(apikey, secret, userkey string) *WykopAPIV1 {

	WykopAPI := gokop.CreateAPIBase(apikey, secret)
	apiClient := &WykopAPIV1{
		WykopAPI: &WykopAPI,
		baseURL:  APIURL,
	}
	return apiClient
}
func (w *WykopAPIV1) request(endpoint string, optionalParams ...gokop.RequestOptionalParam) *gokop.WykopRequest {
	return gokop.CreateRequest(w, w.baseURL, endpoint, optionalParams...)
}
func (w *WykopAPIV1) MakeRequest(endpoint string, target interface{}, optionalParams ...gokop.RequestOptionalParam) error {
	req := w.request(endpoint, optionalParams...)
	data, err := w.SendRequest(req)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	APIErr := models.ErrorResponse{}
	err = gokop.DecodeJSON(data, &APIErr)
	if err != nil {
		return &APIErr
	}
	return gokop.DecodeJSON(data, &target)
}
