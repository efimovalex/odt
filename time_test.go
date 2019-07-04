package odt_test

import (
	"reflect"
	"testing"
	"time"

	date "github.com/efimovalex/odt"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		d       *date.Time
		args    args
		wantErr bool
	}{
		{
			"Success null value",
			&date.Time{},
			args{[]byte(`null`)},
			false,
		},
		{
			"Success date value",
			&date.Time{},
			args{[]byte(`"18:10:02"`)},
			false,
		},
		{
			"Success stripped of quotation marks ",
			&date.Time{},
			args{[]byte(`22:10:12`)},
			false,
		},
		{
			"Error non valid time value",
			&date.Time{},
			args{[]byte(`"18:10"`)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Time.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	tm, _ := time.Parse(time.RFC3339, "0001-01-01T23:59:59+03:00")
	tests := []struct {
		name    string
		d       *date.Time
		want    []byte
		wantErr bool
	}{
		{
			"Success null value",
			&date.Time{},
			[]byte(`null`),
			false,
		},
		{
			"Success good value",
			&date.Time{&tm},
			[]byte(`"23:59:59"`),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}
