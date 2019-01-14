package gokop

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func Test_requestMethod_ToString(t *testing.T) {
	tests := []struct {
		name string
		r    requestMethod
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ToString(); got != tt.want {
				t.Errorf("requestMethod.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeRequest(t *testing.T) {
	type args struct {
		version APIVersionT
	}
	tests := []struct {
		name string
		args args
		want *WykopRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeRequest(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopRequest_IsSigned(t *testing.T) {
	tests := []struct {
		name string
		req  *WykopRequest
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.IsSigned(); got != tt.want {
				t.Errorf("WykopRequest.IsSigned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopRequest_Method(t *testing.T) {
	tests := []struct {
		name string
		req  *WykopRequest
		want requestMethod
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.Method(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopRequest.Method() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopRequest_GetPostParams(t *testing.T) {
	tests := []struct {
		name string
		req  *WykopRequest
		want url.Values
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.GetPostParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopRequest.GetPostParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopRequest_SetPostParams(t *testing.T) {
	type args struct {
		params url.Values
	}
	tests := []struct {
		name string
		req  *WykopRequest
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.req.SetPostParams(tt.args.params)
		})
	}
}

func TestWykopRequest_GetHeaders(t *testing.T) {
	tests := []struct {
		name string
		req  *WykopRequest
		want http.Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.GetHeaders(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopRequest.GetHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
