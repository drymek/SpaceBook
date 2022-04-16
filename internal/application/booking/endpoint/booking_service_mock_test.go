package endpoint_test

import (
	"context"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"github.com/stretchr/testify/mock"
)

type BookingServiceMock struct {
	mock.Mock
}

func (b *BookingServiceMock) Delete(id string) error {
	args := b.Called(id)
	return args.Error(0)
}

func (b *BookingServiceMock) List() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingServiceMock) Create(ctx context.Context, booking *model.Booking) error {
	args := b.Called(ctx, booking)
	return args.Error(0)
}
