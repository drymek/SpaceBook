package repository

import (
	"errors"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

var ErrBookingAlreadyExists = errors.New("booking already exists")

type BookingRepository interface {
	Create(booking *model.Booking) error
}