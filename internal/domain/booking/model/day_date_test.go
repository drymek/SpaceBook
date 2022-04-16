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

func (s *DayDateSuite) TestStart() {
	date, err := NewDayDateFromString("2020-01-01")
	s.NoError(err)
	s.Equal("2020-01-01T00:00:00.000Z", date.Start().Format("2006-01-02T15:04:05.000Z"))
}

func (s *DayDateSuite) TestEnd() {
	date, err := NewDayDateFromString("2020-01-01")
	s.NoError(err)
	s.Equal("2020-01-01T23:59:59.000Z", date.End().Format("2006-01-02T15:04:05.000Z"))
}
