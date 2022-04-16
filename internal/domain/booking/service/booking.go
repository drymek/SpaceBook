package service

import (
	"context"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
	"dryka.pl/SpaceBook/internal/domain/booking/spacex"
)

type BookingService interface {
	Create(booking *model.Booking) error
	List() ([]model.Booking, error)
}

type bookingService struct {
	repository repository.BookingRepository
	client     spacex.SpaceXClient
	timetable  Timetable
}

func (b *bookingService) List() ([]model.Booking, error) {
	return b.repository.List()
}

func (b *bookingService) Create(booking *model.Booking) error {
	ctx := context.Background()
	launches, err := b.client.GetLaunches(ctx, booking.LaunchDate, booking.LaunchpadID)
	if err != nil {
		return ErrBookingService(err)
	}

	if len(launches) > 0 {
		return ErrBookingService(ErrBookingDateConflict)
	}

	timetable := NewTimetable()

	if (*timetable)[booking.LaunchDate.Weekday()] != booking.DestinationID {
		return ErrBookingService(ErrBookingDateTimetable)
	}

	return b.repository.Create(booking)
}

func NewBookingService(repository repository.BookingRepository, client spacex.SpaceXClient) BookingService {
	return &bookingService{
		repository: repository,
		client:     client,
	}
}
