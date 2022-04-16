package model

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BookingSuite struct {
	suite.Suite
	ValidLaunchpadID     LaunchpadID
	InvalidLaunchpadID   LaunchpadID
	InvalidDestinationID DestinationID
	ValidDestinationID   DestinationID
}

func TestSuite(t *testing.T) {
	s := new(BookingSuite)
	s.ValidLaunchpadID = VandenbergSpaceForceBase1
	s.InvalidLaunchpadID = "fake-id"
	s.ValidDestinationID = Mars
	s.InvalidDestinationID = "fake-id"
	suite.Run(t, s)
}

func (s *BookingSuite) TestValidBooking() {
	launchDate, err := NewDayDateFromString("2222-02-01")
	s.NoError(err)
	birthday, err := NewDayDateFromString("1999-02-01")
	s.NoError(err)
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      birthday,
		LaunchpadID:   s.ValidLaunchpadID,
		DestinationID: s.ValidDestinationID,
		LaunchDate:    launchDate,
	}
	err = b.Validate()
	s.NoError(err)
}

func (s *BookingSuite) TestInvalidPast() {
	launchDate, err := NewDayDateFromString("2000-01-17")
	s.NoError(err)
	birthday, err := NewDayDateFromString("1999-02-01")
	s.NoError(err)
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      birthday,
		LaunchpadID:   s.ValidLaunchpadID,
		DestinationID: s.ValidDestinationID,
		LaunchDate:    launchDate,
	}
	err = b.Validate()
	s.Error(err)
	s.True(errors.Is(err, ErrBookingValidation))
	s.Contains(err.Error(), "past launch date")
}

func (s *BookingSuite) TestInvalidBookingLaunchpad() {
	launchDate, err := NewDayDateFromString("2222-02-01")
	s.NoError(err)
	birthday, err := NewDayDateFromString("1999-02-01")
	s.NoError(err)
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      birthday,
		LaunchpadID:   s.InvalidLaunchpadID,
		DestinationID: Mars,
		LaunchDate:    launchDate,
	}

	err = b.Validate()
	s.Error(err)
	s.True(errors.Is(err, ErrBookingValidation))
	s.Contains(err.Error(), "launchpad_id")
}

func (s *BookingSuite) TestInvalidBookingDestination() {
	launchDate, err := NewDayDateFromString("2222-02-01")
	s.NoError(err)
	birthday, err := NewDayDateFromString("1999-02-01")
	s.NoError(err)
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      birthday,
		LaunchpadID:   s.ValidLaunchpadID,
		DestinationID: s.InvalidDestinationID,
		LaunchDate:    launchDate,
	}

	err = b.Validate()
	s.Error(err)
	s.True(errors.Is(err, ErrBookingValidation))
	s.Contains(err.Error(), "destination_id")
}
