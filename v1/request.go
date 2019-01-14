package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cymruu/gokop"
)

type OptionalParamV1 func(*WykopRequestV1)
type MethodParams []string
type APIParam struct {
	Name  string
	Value string
}
type WykopRequestV1 struct {
	*gokop.WykopRequest
	//params
	_APIParams   []APIParam
	methodParams MethodParams
}

func CreateRequest(client *WykopAPIV1, endpoint string, optionalParams ...OptionalParamV1) *WykopRequestV1 {
	req := &WykopRequestV1{
		gokop.InitializeRequest(client.APIVersion()), make([]APIParam, 0), make(MethodParams, 0),
	}
	req.URL = client.APIURL()
	req.Endpoint = endpoint
	if client.Userkey() != "" {
		req._APIParams = append(req._APIParams, APIParam{"userkey", client.Userkey()})
	}
	req._APIParams = append(req._APIParams, APIParam{"appkey", client.APIKey()})
	for _, param := range client.defaultOptions {
		param(req)
	}
	for _, param := range optionalParams {
		param(req)
	}

	if req.Method() == gokop.POST {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("User-Agent", client.Useragent())
	return req
}
func OpMethodParams(params ...string) OptionalParamV1 {
	return func(req *WykopRequestV1) {
		req.methodParams = append(req.methodParams, params...)
	}
}
func OPAPIParams(params ...APIParam) OptionalParamV1 {
	return func(req *WykopRequestV1) {
		req._APIParams = append(req._APIParams, params...)
	}
}
func (req *WykopRequestV1) ToHTTPRequest() (*http.Request, error) {
	request, err := http.NewRequest(req.Method().ToString(), req.BuildURL(), strings.NewReader(req.PostParams.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header = req.Header
	return request, nil
}
func (req *WykopRequestV1) BuildURL() string {
	URL := fmt.Sprintf("%s/%s/", strings.TrimSuffix(req.URL, "/"), req.Endpoint)
	if req.methodParams != nil {
		URL += fmt.Sprintf("%s/", strings.Join(req.methodParams, "/"))
	}
	if req._APIParams != nil {
		for x := range req._APIParams {
			URL += fmt.Sprintf("%s,%s,", req._APIParams[x].Name, req._APIParams[x].Value)
		}
	}
	return strings.TrimSuffix(URL, ",")
}
