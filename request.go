package gokop

import (
	"net/http"
	"net/url"
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

type MethodParam string
type APIParam struct {
	name  string
	value string
}
type RequestOptionalParam func(*WykopRequest)
type IWykopRequest interface {
	IsSigned() bool
	Method() requestMethod

	BuildURL() string
	ToRequest() (*http.Request, error)
	// Send() ([]byte, error)

	GetHeaders() http.Header
	GetPostParams() url.Values
}
type WykopRequest struct {
	_v APIVersionT //verion of api

	URL      string
	Endpoint string

	//shared by both versions
	Header     http.Header
	PostParams url.Values
}

func InitializeRequest() *WykopRequest {
	return &WykopRequest{
		Header:     make(http.Header),
		PostParams: make(url.Values),
	}
}
func SetPostParams(params url.Values) RequestOptionalParam {
	return func(r *WykopRequest) {
		r.PostParams = params
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
func (req *WykopRequest) GetHeaders() http.Header {
	return req.Header
}
