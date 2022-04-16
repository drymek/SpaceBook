package response

import "dryka.pl/SpaceBook/internal/domain/booking/model"

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

type BookingList struct {
	Items []Booking `json:"items"`
	code  int
}

func (b *BookingList) StatusCode() int {
	return b.code
}

func NewBookingResponse(code int, bookings []model.Booking) *BookingList {
	var bookingsResponse []Booking
	for _, booking := range bookings {
		bookingsResponse = append(bookingsResponse, Booking{
			ID:            booking.ID,
			Firstname:     booking.Firstname,
			Lastname:      booking.Lastname,
			Gender:        booking.Gender,
			Birthday:      booking.Birthday.String(),
			LaunchpadID:   string(booking.LaunchpadID),
			DestinationID: string(booking.DestinationID),
			LaunchDate:    booking.LaunchDate.String(),
		})
	}
	return &BookingList{
		Items: bookingsResponse,
		code:  code,
	}
}
