package service

import (
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
)

type BookingService interface {
	Create(booking model.Booking) error
}

type bookingService struct {
	repository repository.BookingRepository
	client     SpaceXClient
	timetable  Timetable
}

func (b *bookingService) Create(booking model.Booking) error {
	launches, err := b.client.GetLaunches(booking.LaunchDate, booking.LaunchpadID)
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

func NewBookingService(repository repository.BookingRepository, client SpaceXClient) BookingService {
	return &bookingService{
		repository: repository,
		client:     client,
	}
}
