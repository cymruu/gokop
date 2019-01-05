package v1

import (
	"crypto/md5"
	"fmt"
	"sort"
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

func CreateRequest(client gokop.IClient, url, endpoint string, optionalParams ...WykopRequestV1OptionalParamF) *WykopRequestV1 {
	req := &WykopRequestV1{
		gokop.InitializeRequest(), make([]APIParam, 0), make(MethodParams, 0),
	}
	if client.Userkey() != "" {
		req._APIParams = append(req._APIParams, APIParam{"userkey", client.Userkey()})
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
func (req *WykopRequestV1) Sign(w *WykopAPIV1) {
	tosign := w.Secret() + req.BuildURL()
	if req.PostParams != nil {
		keys := make([]string, len(req.PostParams))
		for k := range req.PostParams {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, v := range keys {
			tosign += req.PostParams[v][0] + ","
		}
		tosign = tosign[:len(tosign)-1]
	}
	fmt.Println(tosign)
	checksum := md5.Sum([]byte(tosign))
	req.Header.Add("apisign", fmt.Sprintf("%x", checksum))
	fmt.Println(req.Header.Get("apisign"))
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
