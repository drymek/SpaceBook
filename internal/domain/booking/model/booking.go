package model

import (
	"fmt"
)

type Booking struct {
	ID            string        `json:"id"`
	Firstname     string        `json:"firstname"`
	Lastname      string        `json:"lastname"`
	Gender        string        `json:"gender"`
	Birthday      DayDate       `json:"birthday"`
	LaunchpadID   LaunchpadID   `json:"launchpad_id"`
	DestinationID DestinationID `json:"destination_id"`
	LaunchDate    DayDate       `json:"launchDate"`
}

func (b Booking) Validate() error {
	if len(b.Firstname) == 0 {
		return fmt.Errorf("%w: invalid firstname", ErrBookingValidation)
	}

	err := b.validateLaunchpadID()
	if err != nil {
		return err
	}

	err = b.validateDestinationID()
	if err != nil {
		return err
	}

	return nil
}

func (b Booking) validateLaunchpadID() error {
	switch b.LaunchpadID {
	case VandenbergSpaceForceBase1, CapeCanaveral1, BocaChicaVillage, OmelekIsland, VandenbergSpaceForceBase2, CapeCanaveral2:
		return nil
	}

	return fmt.Errorf("%w: invalid launchpad_id", ErrBookingValidation)
}

func (b Booking) validateDestinationID() error {
	switch b.DestinationID {
	case Mars, Moon, Pluto, AsteroidBelt, Europa, Titan, Ganymede:
		return nil
	}

	return fmt.Errorf("%w: invalid destination_id", ErrBookingValidation)
}

func NewBooking(id string, firstname string, lastname string, gender string, birthday DayDate, launchpadID LaunchpadID, destinationID DestinationID, launchDate DayDate) (Booking, error) {
	booking := Booking{
		ID:            id,
		Firstname:     firstname,
		Lastname:      lastname,
		Gender:        gender,
		Birthday:      birthday,
		LaunchpadID:   launchpadID,
		DestinationID: destinationID,
		LaunchDate:    launchDate,
	}

	err := booking.Validate()
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}
