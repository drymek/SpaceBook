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
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "",
		LaunchpadID:   s.ValidLaunchpadID,
		DestinationID: s.ValidDestinationID,
		LaunchDate:    "",
	}

	err := b.Validate()
	s.NoError(err)
}

func (s *BookingSuite) TestInvalidBookingLaunchpad() {
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "",
		LaunchpadID:   s.InvalidLaunchpadID,
		DestinationID: Mars,
		LaunchDate:    "",
	}

	err := b.Validate()
	s.Error(err)
	s.True(errors.Is(err, ErrBookingValidation))
	s.Contains(err.Error(), "launchpad_id")
}

func (s *BookingSuite) TestInvalidBookingDestination() {
	b := Booking{
		ID:            "",
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "",
		LaunchpadID:   s.ValidLaunchpadID,
		DestinationID: s.InvalidDestinationID,
		LaunchDate:    "",
	}

	err := b.Validate()
	s.Error(err)
	s.True(errors.Is(err, ErrBookingValidation))
	s.Contains(err.Error(), "destination_id")
}
