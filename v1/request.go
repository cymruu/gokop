package v1

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cymruu/gokop"
)

type MethodParams []string
type APIParam struct {
	name  string
	value string
}
type WykopRequestV1OptionalParamF func(*WykopRequestV1)
type WykopRequestV1 struct {
	*gokop.WykopRequest
	//params
	_APIParams   []APIParam
	methodParams MethodParams
}

func CreateRequest(client gokop.IClient, endpoint string, optionalParams ...WykopRequestV1OptionalParamF) *WykopRequestV1 {
	req := &WykopRequestV1{
		gokop.InitializeRequest(), make([]APIParam, 0), make(MethodParams, 0),
	}
	req.URL = client.APIURL()
	req.Endpoint = endpoint
	if client.Userkey() != "" {
		req._APIParams = append(req._APIParams, APIParam{"userkey", client.Userkey()})
	}
	req._APIParams = append(req._APIParams, APIParam{"appkey", client.APIKey()})
	for _, param := range optionalParams {
		param(req)
	}

	if req.Method() == gokop.POST {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	req.Header.Add("User-Agent", client.Useragent())
	return req
}
func AddMethodParams(params MethodParams) WykopRequestV1OptionalParamF {
	return func(r *WykopRequestV1) {
		r.methodParams = append(r.methodParams, params...)
	}
}
func AddAPIParams(params ...APIParam) WykopRequestV1OptionalParamF {
	return func(r *WykopRequestV1) {
		r._APIParams = append(r._APIParams, params...)
	}
}
func (req *WykopRequestV1) ToRequest() (*http.Request, error) {
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
			URL += fmt.Sprintf("%s,%s,", req._APIParams[x].name, req._APIParams[x].value)
		}
	}
	return strings.TrimSuffix(URL, ",")
}
