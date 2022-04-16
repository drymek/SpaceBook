package service

import (
	"context"
	"fmt"
	"testing"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookingSuite struct {
	suite.Suite
}

func TestDayDateSuite(t *testing.T) {
	s := new(BookingSuite)

	suite.Run(t, s)
}

type SpaceXClientMock struct {
	mock.Mock
}

func (s *SpaceXClientMock) GetLaunches(ctx context.Context, date model.DayDate, id model.LaunchpadID) ([]string, error) {
	args := s.Called(date, id)
	return args.Get(0).([]string), args.Error(1)
}

func (s *BookingSuite) TestPersist() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	booking := &model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   model.VandenbergSpaceForceBase1,
		DestinationID: model.Pluto,
		LaunchDate:    date,
	}
	repository := new(BookingRepositoryMock)
	repository.On("Create", booking).Return(nil)

	spacexClient := new(SpaceXClientMock)
	spacexClient.On("GetLaunches", date, model.VandenbergSpaceForceBase1).Return([]string{}, nil)
	err = NewBookingService(repository, spacexClient).Create(booking)
	s.NoError(err)
	repository.AssertCalled(s.T(), "Create", booking)
}

func (s *BookingSuite) TestSpaceXBookingExists() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	booking := &model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   model.VandenbergSpaceForceBase1,
		DestinationID: model.Pluto,
		LaunchDate:    date,
	}
	repository := new(BookingRepositoryMock)
	repository.On("Create", booking).Return(nil)

	spacexClient := new(SpaceXClientMock)
	spacexClient.On("GetLaunches", date, model.VandenbergSpaceForceBase1).Return([]string{"3a50ae198086"}, nil)
	err = NewBookingService(repository, spacexClient).Create(booking)
	s.Error(err)
	s.ErrorIs(err, ErrBookingDateConflict)
	repository.AssertNotCalled(s.T(), "Create", mock.Anything)
}

func (s *BookingSuite) TestSpaceXBookingClientError() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	booking := &model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   model.VandenbergSpaceForceBase1,
		DestinationID: model.Pluto,
		LaunchDate:    date,
	}
	repository := new(BookingRepositoryMock)
	repository.On("Create", booking).Return(nil)

	spacexClient := new(SpaceXClientMock)
	ErrSpacex := fmt.Errorf("SpaceX error")
	spacexClient.On("GetLaunches", date, model.VandenbergSpaceForceBase1).Return([]string{}, ErrSpacex)
	err = NewBookingService(repository, spacexClient).Create(booking)
	s.Error(err)
	s.ErrorIs(err, ErrSpacex)
	repository.AssertNotCalled(s.T(), "Create", mock.Anything)
}

func (s *BookingSuite) TestWrongDestination() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	booking := &model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   model.VandenbergSpaceForceBase1,
		DestinationID: model.Moon,
		LaunchDate:    date,
	}
	repository := new(BookingRepositoryMock)
	repository.On("Create", booking).Return(nil)

	spacexClient := new(SpaceXClientMock)
	spacexClient.On("GetLaunches", date, model.VandenbergSpaceForceBase1).Return([]string{}, nil)
	err = NewBookingService(repository, spacexClient).Create(booking)
	s.Error(err)
	s.ErrorIs(err, ErrBookingDateTimetable)
	repository.AssertNotCalled(s.T(), "Create", mock.Anything)
}

func (s *BookingSuite) TestListError() {
	repository := new(BookingRepositoryMock)
	ErrListing := fmt.Errorf("Listing error")
	repository.On("List").Return([]model.Booking{}, ErrListing)
	spacexClient := new(SpaceXClientMock)

	_, err := NewBookingService(repository, spacexClient).List()
	s.Error(err)
	s.ErrorIs(err, ErrListing)
}

func (s *BookingSuite) TestListSuccess() {
	repository := new(BookingRepositoryMock)
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	list := []model.Booking{
		{
			ID:            "123",
			Firstname:     "Marcin",
			Lastname:      "Dryka",
			Gender:        "Male",
			Birthday:      date,
			LaunchpadID:   model.VandenbergSpaceForceBase1,
			DestinationID: model.Moon,
			LaunchDate:    model.DayDate{},
		},
	}
	repository.On("List").Return(list, nil)
	spacexClient := new(SpaceXClientMock)

	got, err := NewBookingService(repository, spacexClient).List()
	s.NoError(err)
	s.Equal(got, list)
}

type BookingRepositoryMock struct {
	mock.Mock
}

func (b *BookingRepositoryMock) Find(s string) (model.Booking, error) {
	args := b.Called(s)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BookingRepositoryMock) Delete(s string) error {
	args := b.Called(s)
	return args.Error(0)
}

func (b *BookingRepositoryMock) List() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BookingRepositoryMock) Create(booking *model.Booking) error {
	args := b.Called(booking)
	return args.Error(0)
}
