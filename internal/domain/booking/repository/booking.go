package repository

import (
	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type BookingRepository interface {
	Create(booking model.Booking) error
}
