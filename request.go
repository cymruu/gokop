package gokop

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type requestMethod uint8

func (r requestMethod) ToString() string {
	names := []string{"GET", "POST"}
	return names[r]
}

const (
	GET  requestMethod = 0
	POST requestMethod = 1
)

type methodParams []string
type apiParam struct {
	name  string
	value string
}
type RequestOptionalParam func(*WykopRequest)
type WykopRequest struct {
	_v string //verion of api

	URL      string
	Endpoint string

	methodParams methodParams
	apiParams    []apiParam
	postParams   url.Values
	Header       http.Header
}

func CreateRequest(client IClient, url, endpoint string, optionalParams ...RequestOptionalParam) *WykopRequest {
	req := &WykopRequest{
		_v:       client.APIVersion(),
		URL:      url,
		Endpoint: endpoint,

		Header:       make(http.Header),
		methodParams: make(methodParams, 0),
		apiParams:    make([]apiParam, 0),
	}
	req.apiParams = append(req.apiParams, apiParam{"appkey", client.APIKey()})
	if client.Userkey() != "" {
		req.apiParams = append(req.apiParams, apiParam{"userkey", client.Userkey()})
	}
	for _, param := range optionalParams {
		param(req)
	}
	if req.Method() == POST {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Header.Add("User-Agent", client.Useragent())
	return req
}
func (req *WykopRequest) IsSigned() bool {
	return req.Header.Get("apisign") != ""
}
func AddMethodParams(params methodParams) RequestOptionalParam {
	return func(r *WykopRequest) {
		r.methodParams = append(r.methodParams, params...)
	}
}
func AddAPIParams(params []apiParam) RequestOptionalParam {
	return func(r *WykopRequest) {
		r.apiParams = append(r.apiParams, params...)
	}
}
func SetPostParams(params url.Values) RequestOptionalParam {
	return func(r *WykopRequest) {
		r.postParams = params
	}
}
func (req *WykopRequest) Method() requestMethod {
	if req.postParams == nil {
		return POST
	}
	return GET
}
func (req *WykopRequest) BuildURL() string {
	URL := fmt.Sprintf("%s/%s/", strings.TrimSuffix(req.URL, "/"), req.Endpoint)
	if req.methodParams != nil {
		URL += fmt.Sprintf("%s/", strings.Join(req.methodParams, "/"))
	}
	if req.apiParams != nil {
		for x := range req.apiParams {
			URL += fmt.Sprintf("%s,%s,", req.apiParams[x].name, req.apiParams[x].value)
		}
	}
	return strings.TrimSuffix(URL, ",")
}
