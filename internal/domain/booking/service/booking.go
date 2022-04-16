package service

import (
	"context"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
	"dryka.pl/SpaceBook/internal/domain/booking/spacex"
	"github.com/google/uuid"
)

type BookingService interface {
	Create(ctx context.Context, booking *model.Booking) error
	List() ([]model.Booking, error)
	Delete(id string) error
}

type bookingService struct {
	repository repository.BookingRepository
	client     spacex.SpaceXClient
	timetable  Timetable
}

func (b *bookingService) Delete(id string) error {
	return b.repository.Delete(id)
}

func (b *bookingService) List() ([]model.Booking, error) {
	return b.repository.List()
}

func (b *bookingService) Create(ctx context.Context, booking *model.Booking) error {
	err := booking.Validate()
	if err != nil {
		return err
	}

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

	if booking.ID == "" {
		booking.ID = b.GenerateID()
	}
	return b.repository.Create(booking)
}

func (b *bookingService) GenerateID() string {
	return uuid.New().String()
}

func NewBookingService(repository repository.BookingRepository, client spacex.SpaceXClient) BookingService {
	return &bookingService{
		repository: repository,
		client:     client,
	}
}
