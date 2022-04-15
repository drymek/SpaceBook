package model

type Booking struct {
	ID            string `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadID   string `json:"launchpadID"`
	DestinationID string `json:"destinationID"`
	LaunchDate    string `json:"launchDate"`
}

func (b Booking) Validate() error {
	if len(b.Firstname) == 0 {
		return ErrInvalidFirstname
	}

	return nil
}

func NewBooking(firstname string, lastname string, gender string, birthday string, launchpadID string, destinationID string, launchDate string) (Booking, error) {
	booking := Booking{
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
