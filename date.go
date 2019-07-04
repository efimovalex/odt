package odt

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const DefaultLayout = "2006-01-02"

var dateLayout = DefaultLayout

type Date struct {
	*time.Time
}

func SetDateFormat(f string) {
	dateLayout = f
}

func GetDateFormat() string {
	return dateLayout
}

func NewDate(t time.Time) *Date {
	return &Date{&t}
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = nil

		return
	}
	t, err := time.Parse(dateLayout, s)

	d.Time = &t

	return
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(dateLayout))), nil
}

func (d Date) Value() (driver.Value, error) {
	if d.Time == nil {
		return driver.Value(nil), nil
	}

	return driver.Value(*d.Time), nil
}

// Scan - Implement the database/sql scanner interface
func (d *Date) Scan(value interface{}) error {
	// if value is nil
	if value == nil {
		*d = Date{nil}
		return nil
	}
	if bv, err := convertDateValue(value); err == nil {
		// if this is a time type
		if v, ok := bv.(time.Time); ok {
			*d = Date{&v}
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan date.Date")
}

func convertDateValue(src interface{}) (driver.Value, error) {
	switch s := src.(type) {
	case time.Time:
		return s, nil
	case string:
		tm, err := time.Parse(DefaultLayout, s)
		if err != nil {
			return nil, fmt.Errorf("sql/driver: couldn't convert %q into type Date", s)
		}
		return tm, nil
	case []byte:
		tm, err := time.Parse(DefaultLayout, string(s))
		if err != nil {
			return nil, fmt.Errorf("sql/driver: couldn't convert %q into type Date", s)
		}
		return tm, nil
	}

	return nil, fmt.Errorf("sql/driver: couldn't convert %v (%T) into type Date", src, src)
}
