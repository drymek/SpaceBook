package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type DayDate struct {
	time.Time
}

var format = "2006-01-02"

func (d *DayDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("%w: failed to unmarshal date", ErrDayDateValidation)
	}

	t, err := time.Parse(format, s)
	if err != nil {
		return fmt.Errorf("%w: failed to parse date", ErrDayDateValidation)
	}

	d.Time = t
	return nil
}

func (d *DayDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DayDate) String() string {
	return d.Time.Format(format)
}

func NewDayDateFromString(value string) (DayDate, error) {
	d, err := time.Parse(format, value)
	if err != nil {
		return DayDate{}, fmt.Errorf("%w: failed to parse date", ErrDayDateValidation)
	}

	return DayDate{
		Time: d,
	}, nil
}
