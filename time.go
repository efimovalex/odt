package odt

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const DefaultTimeLayout = "15:04:05"

var timeLayout = DefaultTimeLayout

type Time struct {
	*time.Time
}

func NewTime(t time.Time) *Time {
	return &Time{&t}
}

func SetTimeFormat(f string) {
	timeLayout = f
}

func GetTimeFormat() string {
	return timeLayout
}

func ParseTimeFromBytes(b []byte) (*Time, error) {
	t, err := time.Parse(timeLayout, string(b))

	return &Time{&t}, err
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		t.Time = nil

		return
	}
	tm, err := time.Parse(timeLayout, s)

	t.Time = &tm

	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(timeLayout))), nil
}

// Scan - Implement the database/sql scanner interface
func (t *Time) Scan(value interface{}) error {
	// if value is nil
	if value == nil {
		*t = Time{nil}
		return nil
	}
	if bv, err := convertTimeValue(value); err == nil {
		// if this is a time type
		if v, ok := bv.(time.Time); ok {
			*t = Time{&v}
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan date.Time")
}

func (t Time) Value() (driver.Value, error) {
	if t.Time == nil {
		return driver.Value(nil), nil
	}

	return driver.Value(*t.Time), nil
}

func convertTimeValue(src interface{}) (driver.Value, error) {
	switch s := src.(type) {
	case string:
		tm, err := time.Parse(DefaultTimeLayout, s)
		if err != nil {
			return nil, fmt.Errorf("sql/driver: couldn't convert %q into type Time", s)
		}
		return tm, nil
	case []byte:
		tm, err := time.Parse(DefaultTimeLayout, string(s))
		if err != nil {
			return nil, fmt.Errorf("sql/driver: couldn't convert %q into type Time", s)
		}
		return tm, nil
	}

	return nil, fmt.Errorf("sql/driver: couldn't convert %v (%T) into type Time", src, src)
}
