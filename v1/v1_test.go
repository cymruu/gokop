package v1

import (
	"reflect"
	"testing"

	"github.com/cymruu/gokop"
)

var client = CreateWykopV1API("apikey", "secret", "")

func TestWykopAPIV1_AddDefaultParameters(t *testing.T) {
	type args struct {
		params []OptionalParamV1
	}
	tests := []struct {
		name string
		w    *WykopAPIV1
		args args
	}{
		{
			name: "add default parameter",
			w:    client,
			args: args{params: []OptionalParamV1{OPAPIParams(APIParam{"page", "5"})}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 1; i < 3; i++ {
				tt.w.AddDefaultParameters(tt.args.params...)
				if len(tt.w.defaultOptions) != i {
					t.Error("Failed to add default parameter")
				}
			}
		})
	}
}

func TestWykopAPIV1_Useragent(t *testing.T) {
	tests := []struct {
		name string
		w    *WykopAPIV1
		want string
	}{
		{
			name: "Client without useragent",
			w:    client,
			want: gokop.DefaultUseragent,
		}, {
			name: "Client with useragent set",
			w:    &WykopAPIV1{useragent: "testing"},
			want: "testing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Useragent(); got != tt.want {
				t.Errorf("WykopAPIV1.Useragent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_SetUseragent(t *testing.T) {
	type args struct {
		useragent string
	}
	tests := []struct {
		name string
		w    *WykopAPIV1
		args args
		want string
	}{
		{
			name: "Client without useragent",
			w:    client,
			args: args{useragent: ""},
			want: gokop.DefaultUseragent,
		}, {
			name: "Client with useragent set",
			w:    &WykopAPIV1{},
			args: args{useragent: "testing"},
			want: "testing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.SetUseragent(tt.args.useragent)
			if got := tt.w.Useragent(); got != tt.want {
				t.Errorf("WykopAPIV1.SetUseragent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_APIURL(t *testing.T) {
	tests := []struct {
		name string
		w    *WykopAPIV1
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.APIURL(); got != tt.want {
				t.Errorf("WykopAPIV1.APIURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_APIVersion(t *testing.T) {
	tests := []struct {
		name string
		w    *WykopAPIV1
		want gokop.APIVersionT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.APIVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopAPIV1.APIVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_Userkey(t *testing.T) {
	tests := []struct {
		name string
		w    *WykopAPIV1
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Userkey(); got != tt.want {
				t.Errorf("WykopAPIV1.Userkey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_SetUserkey(t *testing.T) {
	client := CreateWykopV1API("", "", "")
	type args struct {
		userkey string
	}
	tests := []struct {
		name string
		w    *WykopAPIV1
		args args
		want string
	}{
		{name: "Set userkey to abcde", w: client, args: args{userkey: "abcde"}, want: "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.SetUserkey(tt.args.userkey)
			if got := tt.w.Userkey(); got != tt.want {
				t.Errorf("WykopAPIV1.SetUserkey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateWykopV1API(t *testing.T) {
	type args struct {
		apikey  string
		secret  string
		userkey string
	}
	tests := []struct {
		name string
		args args
		want *WykopAPIV1
	}{
		{
			name: "Create Wykop API v1 API Client Without Userkey",
			args: args{
				apikey:  "apikey",
				secret:  "secret",
				userkey: "",
			},
			want: &WykopAPIV1{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWykopV1API(tt.args.apikey, tt.args.secret, tt.args.userkey)
			if got.APIKey() != tt.args.apikey {
				t.Errorf("CreateAPIBase() = %v, want %v", got.APIKey(), tt.args.apikey)
			}
			if got.Secret() != tt.args.secret {
				t.Errorf("CreateAPIBase() = %v, want %v", got.Secret(), tt.args.secret)
			}
			if got.Userkey() != tt.args.userkey {
				t.Errorf("CreateAPIBase() = %v, want %v", got.Userkey(), tt.args.secret)
			}
		})
	}
}

func TestWykopAPIV1_request(t *testing.T) {
	type args struct {
		endpoint       string
		optionalParams []OptionalParamV1
	}
	tests := []struct {
		name string
		w    *WykopAPIV1
		args args
		want *WykopRequestV1
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.request(tt.args.endpoint, tt.args.optionalParams...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopAPIV1.request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_MakeRequest(t *testing.T) {
	type args struct {
		req    *WykopRequestV1
		target interface{}
	}
	tests := []struct {
		name    string
		w       *WykopAPIV1
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.MakeRequest(tt.args.req, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("WykopAPIV1.MakeRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
