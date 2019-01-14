package v2

import (
	"fmt"

	"github.com/cymruu/gokop"
	"github.com/cymruu/gokop/v2/models"
)

const (
	apiVersion = gokop.APIVersionV2
	APIURL     = "https://a2.wykop.pl"
	//response types
	responseJSON  = "json"
	responseJSONP = "jsonp"
	responseXML   = "xml"

	//response filters
	outputCLEAR = "clear"
)

// type ErrorHandlerF func(*models.ErrorResponse, *gokop.WykopRequest)
type WykopAPIV2 struct {
	*gokop.WykopAPI

	userkey string

	useragent string
	baseURL   string

	// errorHandlers map[int]ErrorHandlerF
}

func (w *WykopAPIV2) Useragent() string {
	userAgent := gokop.DefaultUseragent
	if w.userkey != "" {
		userAgent = w.useragent
	}
	return userAgent
}
func (w *WykopAPIV2) SetUseragent(useragent string) {
	w.useragent = useragent
}
func (w *WykopAPIV2) APIURL() string {
	return APIURL
}
func (w *WykopAPIV2) APIVersion() gokop.APIVersionT {
	return apiVersion
}
func (w *WykopAPIV2) Userkey() string {
	return w.userkey
}
func (w *WykopAPIV2) SetUserkey(userkey string) {
	w.userkey = userkey
}

func CreateWykopV2API(apikey, secret, userkey string) *WykopAPIV2 {

	WykopAPI := gokop.CreateAPIBase(apikey, secret)
	apiClient := &WykopAPIV2{
		WykopAPI: &WykopAPI,
		baseURL:  APIURL,
		// errorHandlers: make(map[int]ErrorHandlerF),
	}
	return apiClient
}
func (w *WykopAPIV2) request(endpoint string, optionalParams ...OptionalParamV2) *WykopRequestV2 {
	return CreateRequest(w, endpoint, optionalParams...)
}
func (w *WykopAPIV2) MakeRequest(req *WykopRequestV2, target interface{}) error {
	data, err := w.SendRequest(req)
	fmt.Println(string(data))
	if err != nil {
		return err
	}
	APIResponse := models.Response{}
	err = gokop.DecodeJSON(data, &APIResponse)
	if err != nil {
		return err
	}
	if APIResponse.Error != nil {
		return APIResponse.Error
	}
	return gokop.DecodeJSON(APIResponse.Data, &target)
}
func (w *WykopAPIV2) Entry(id string, params ...OptionalParamV2) (*models.Entry, error) {
	params = append(params, OpAPIParams(id))
	req := w.request("entries/entry", params...)
	entry := new(models.Entry)
	err := w.MakeRequest(req, &entry)
	return entry, err
}
