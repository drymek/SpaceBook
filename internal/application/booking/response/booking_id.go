package response

import "dryka.pl/SpaceBook/internal/domain/booking/model"

type BookingId struct {
	ID   string `json:"id"`
	code int
}

func (b *BookingId) StatusCode() int {
	return b.code
}

func NewBookingIdResponse(statusCode int, booking model.Booking) BookingId {
	return BookingId{
		ID:   booking.ID,
		code: statusCode,
	}
}
