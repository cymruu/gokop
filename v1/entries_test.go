package v1

import (
	"reflect"
	"testing"

	"github.com/cymruu/gokop/v1/models"
)

func TestWykopAPIV1_Index(t *testing.T) {
	type args struct {
		entryID int64
		params  []OptionalParamV1
	}
	tests := []struct {
		name    string
		w       *WykopAPIV1
		args    args
		want    *models.Entry
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.w.Index(tt.args.entryID, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("WykopAPIV1.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopAPIV1.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWykopAPIV1_AddComment(t *testing.T) {
	type args struct {
		entry *models.Entry
		body  string
	}
	tests := []struct {
		name    string
		w       *WykopAPIV1
		args    args
		want    *models.OK_ID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.w.AddComment(tt.args.entry, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("WykopAPIV1.AddComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WykopAPIV1.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
