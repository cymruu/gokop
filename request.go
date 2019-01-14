package gokop

import (
	"net/http"
	"net/url"
)

type requestMethod uint8
type OptionalParamF func(interface{})

func (r requestMethod) ToString() string {
	names := []string{"GET", "POST"}
	return names[r]
}

const (
	GET  requestMethod = 0
	POST requestMethod = 1
)

type APIParam struct {
	name  string
	value string
}
type IWykopRequest interface {
	IsSigned() bool
	Method() requestMethod

	BuildURL() string
	ToHTTPRequest() (*http.Request, error)

	GetHeaders() http.Header

	GetPostParams() url.Values
	SetPostParams(url.Values)
}
type WykopRequest struct {
	_v APIVersionT //verion of api

	URL      string
	Endpoint string

	//shared by both versions
	Header     http.Header
	PostParams url.Values
}

func InitializeRequest(version APIVersionT) *WykopRequest {
	return &WykopRequest{
		_v:         version,
		Header:     make(http.Header),
		PostParams: make(url.Values),
	}
}
func (req *WykopRequest) IsSigned() bool {
	return req.Header.Get("apisign") != ""
}
func (req *WykopRequest) Method() requestMethod {
	if req.PostParams == nil {
		return POST
	}
	return GET
}
func (req *WykopRequest) GetPostParams() url.Values {
	return req.PostParams
}
func (req *WykopRequest) SetPostParams(params url.Values) {
	req.PostParams = params
}
func (req *WykopRequest) GetHeaders() http.Header {
	return req.Header
}
