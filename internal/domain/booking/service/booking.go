package service

import "dryka.pl/SpaceBook/internal/domain/booking/model"

type BookingService interface {
	Create(booking model.Booking) error
}

type bookingService struct {
}

func (b *bookingService) Create(booking model.Booking) error {
	return nil
}

func NewBookingService() BookingService {
	return &bookingService{}
}
