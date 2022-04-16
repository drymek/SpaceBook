package repository

import (
	"sync"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
)

type bookingRepository struct {
	mu         sync.Mutex
	collection map[string]*model.Booking
}

func (b *bookingRepository) List() ([]model.Booking, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	bookings := make([]model.Booking, 0, len(b.collection))
	for _, booking := range b.collection {
		bookings = append(bookings, *booking)
	}

	return bookings, nil
}

func (b *bookingRepository) Create(booking *model.Booking) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.collection[booking.ID]; ok {
		return repository.ErrBookingAlreadyExists
	}

	b.collection[booking.ID] = booking

	return nil
}

func NewBookingRepository() repository.BookingRepository {
	return &bookingRepository{
		mu:         sync.Mutex{},
		collection: make(map[string]*model.Booking),
	}
}
