package odt_test

import (
	"reflect"
	"testing"
	"time"

	date "github.com/efimovalex/odt"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		d       *date.Date
		args    args
		wantErr bool
	}{
		{
			"Success null value",
			&date.Date{},
			args{[]byte(`null`)},
			false,
		},
		{
			"Success date value",
			&date.Date{},
			args{[]byte(`"2018-10-02"`)},
			false,
		},
		{
			"Success stripped of quotation marks ",
			&date.Date{},
			args{[]byte(`2018-10-11`)},
			false,
		},
		{
			"Error non valid date value",
			&date.Date{},
			args{[]byte(`"2018-10"`)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	tm, _ := time.Parse(time.RFC3339, "2019-12-27T23:59:59+03:00")
	tests := []struct {
		name    string
		d       *date.Date
		want    []byte
		wantErr bool
	}{
		{
			"Success null value",
			&date.Date{},
			[]byte(`null`),
			false,
		},
		{
			"Success good value",
			&date.Date{&tm},
			[]byte(`"2019-12-27"`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Date.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
