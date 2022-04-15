package endpoint_test

import (
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"github.com/stretchr/testify/mock"
)

type BookingServiceMock struct {
	mock.Mock
}

func (b *BookingServiceMock) List() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingServiceMock) Create(booking *model.Booking) error {
	args := b.Called(booking)
	return args.Error(0)
}
