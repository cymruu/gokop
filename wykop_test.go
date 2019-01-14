package gokop

import (
	"testing"
)

func TestCreateAPIBase(t *testing.T) {
	type args struct {
		apikey string
		secret string
	}
	tests := []struct {
		name string
		args args
		want WykopAPI
	}{
		{
			name: "Example wykop API Client",
			args: args{
				apikey: "apikey",
				secret: "batman",
			},
			want: WykopAPI{apikey: "apikey", secret: "batman"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateAPIBase(tt.args.apikey, tt.args.secret)
			if got.apikey != tt.args.apikey {
				t.Errorf("CreateAPIBase() = %v, want %v", got.apikey, tt.args.apikey)
			}
			if got.secret != tt.args.secret {
				t.Errorf("CreateAPIBase() = %v, want %v", got.secret, tt.args.secret)
			}
		})
	}
}
func TestWykopAPI_Sign(t *testing.T) {
	//run other test
}
