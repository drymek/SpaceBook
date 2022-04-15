package model

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DayDateSuite struct {
	suite.Suite
}

func TestDayDateSuite(t *testing.T) {
	s := new(DayDateSuite)

	suite.Run(t, s)
}

func (s *DayDateSuite) TestNewDayDateFromStringErr() {
	_, err := NewDayDateFromString("")
	s.Error(err)
}

func (s *DayDateSuite) TestNewDayDateFromString() {
	dayDate, err := NewDayDateFromString("2020-01-01")
	s.NoError(err)
	s.Equal("2020-01-01", dayDate.Format("2006-01-02"))
}

func (s *DayDateSuite) TestMarshal() {
	dayDate, err := NewDayDateFromString("2020-01-01")
	s.NoError(err)
	json, err := dayDate.MarshalJSON()
	s.NoError(err)
	s.Equal(`"2020-01-01"`, string(json))
}

func (s *DayDateSuite) TestUnmarshal() {
	dayDate := &DayDate{}
	err := dayDate.UnmarshalJSON([]byte(`"2020-01-01"`))
	s.NoError(err)
	s.Equal("2020-01-01", dayDate.Format("2006-01-02"))
}

func (s *DayDateSuite) TestUnmarshalError() {
	dayDate := &DayDate{}
	err := dayDate.UnmarshalJSON([]byte(``))
	s.Error(err)
}

func (s *DayDateSuite) TestUnmarshalParseError() {
	dayDate := &DayDate{}
	err := dayDate.UnmarshalJSON([]byte(`{}`))
	s.Error(err)
}
