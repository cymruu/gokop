package v1

import (
	"fmt"

	"github.com/cymruu/gokop"
	"github.com/cymruu/gokop/v1/models"
)

const (
	apiVersion = gokop.APIVersionV1
	APIURL     = "https://a.wykop.pl"
	//response types
	responseJSON  = "json"
	responseJSONP = "jsonp"
	responseXML   = "xml"

	//response filters
	outputCLEAR = "clear"
)

type ErrorHandlerF func(*models.ErrorResponse, *gokop.WykopRequest)
type WykopAPIV1 struct {
	*gokop.WykopAPI

	userkey string

	useragent string
	baseURL   string

	errorHandlers map[int]ErrorHandlerF
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
func (w *WykopAPIV1) APIVersion() gokop.APIVersionT {
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
		WykopAPI:      &WykopAPI,
		baseURL:       APIURL,
		errorHandlers: make(map[int]ErrorHandlerF),
	}
	return apiClient
}
func (w *WykopAPIV1) request(endpoint string, optionalParams ...WykopRequestV1OptionalParamF) *WykopRequestV1 {
	return CreateRequest(w, w.baseURL, endpoint, optionalParams...)
}
func (w *WykopAPIV1) MakeRequest(endpoint string, target interface{}, optionalParams ...WykopRequestV1OptionalParamF) error {
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
