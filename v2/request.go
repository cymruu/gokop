package v2

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cymruu/gokop"
)

type OptionalParamV2 func(*WykopRequestV2)
type APIParams []string
type NamedParam struct {
	Name  string
	Value string
}
type WykopRequestV2 struct {
	*gokop.WykopRequest
	//params
	_APIParams  APIParams
	namedParams []NamedParam
}

func CreateRequest(client *WykopAPIV2, endpoint string, optionalParams ...OptionalParamV2) *WykopRequestV2 {
	req := &WykopRequestV2{
		gokop.InitializeRequest(client.APIVersion()), make(APIParams, 0), make([]NamedParam, 0),
	}
	req.URL = client.APIURL()
	req.Endpoint = endpoint
	if client.Userkey() != "" {
		req.namedParams = append(req.namedParams, NamedParam{"userkey", client.Userkey()})
	}
	req.namedParams = append(req.namedParams, NamedParam{"appkey", client.APIKey()})
	for _, param := range optionalParams {
		param(req)
	}

	if req.Method() == gokop.POST {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("User-Agent", client.Useragent())
	return req
}
func OpAPIParams(params ...string) OptionalParamV2 {
	return func(req *WykopRequestV2) {
		req._APIParams = append(req._APIParams, params...)
	}
}
func OpNamedParams(params ...NamedParam) OptionalParamV2 {
	return func(req *WykopRequestV2) {
		req.namedParams = append(req.namedParams, params...)
	}
}
func (req *WykopRequestV2) ToHTTPRequest() (*http.Request, error) {
	request, err := http.NewRequest(req.Method().ToString(), req.BuildURL(), strings.NewReader(req.PostParams.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header = req.Header
	return request, nil
}
func (req *WykopRequestV2) BuildURL() string {
	URL := fmt.Sprintf("%s/%s/", strings.TrimSuffix(req.URL, "/"), req.Endpoint)
	if req._APIParams != nil {
		URL += fmt.Sprintf("%s/", strings.Join(req._APIParams, "/"))
	}
	if req.namedParams != nil {
		for x := range req._APIParams {
			URL += fmt.Sprintf("%s/%s/", req.namedParams[x].Name, req.namedParams[x].Value)
		}
	}
	return strings.TrimSuffix(URL, ",")
}
