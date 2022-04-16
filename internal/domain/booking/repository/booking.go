package repository

import (
	"errors"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

var ErrBookingAlreadyExists = errors.New("booking already exists")

type BookingRepository interface {
	Create(booking *model.Booking) error
	List() ([]model.Booking, error)
	Find(string) (model.Booking, error)
	Delete(string) error
}
