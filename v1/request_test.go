package v1

import (
	"testing"
)

func TestCreateRequest(t *testing.T) {
	type args struct {
		client         *WykopAPIV1
		endpoint       string
		optionalParams []OptionalParamV1
	}
	tests := []struct {
		name string
		args args
		want *WykopRequestV1
	}{
		{
			name: "Entry index",
			args: args{
				client:         client,
				endpoint:       "entry/index",
				optionalParams: []OptionalParamV1{OpMethodParams("1")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateRequest(tt.args.client, tt.args.endpoint, tt.args.optionalParams...)
			got.URL = APIURL
			if got.URL != APIURL {
				t.Errorf("CreateRequest() url = %v, want %v", got, tt.want)
			}
			if got.Endpoint != tt.args.endpoint {
				t.Errorf("CreateRequest() endpoint = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestWykopRequestV1_BuildURL(t *testing.T) {
	tests := []struct {
		name string
		req  *WykopRequestV1
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.req.BuildURL(); got != tt.want {
				t.Errorf("WykopRequestV1.BuildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
