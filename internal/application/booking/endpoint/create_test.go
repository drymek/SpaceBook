package endpoint_test

import (
	"context"
	"net/http"
	"testing"

	"dryka.pl/SpaceBook/internal/application/booking/endpoint"
	"dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	service2 "dryka.pl/SpaceBook/internal/domain/booking/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EndpointSuite struct {
	suite.Suite
}

func TestCreateSuite(t *testing.T) {
	suite.Run(t, new(EndpointSuite))
}

func (s *EndpointSuite) TestHandleErrorFromService() {
	req := request.BookingRequest{
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "2020-01-01",
		LaunchpadID:   string(model.VandenbergSpaceForceBase1),
		DestinationID: model.Moon,
		LaunchDate:    "2020-01-01",
	}

	serviceErr := service2.ErrBookingService(service2.ErrBookingDateConflict)
	service := new(BookingServiceMock)
	service.On("Create", mock.Anything).
		Return(serviceErr)

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, serviceErr)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *EndpointSuite) TestHandleErrorFromBooking() {
	req := request.BookingRequest{
		Firstname:     "",
		Lastname:      "",
		Gender:        "",
		Birthday:      "2020-01-01",
		LaunchpadID:   string(model.VandenbergSpaceForceBase1),
		DestinationID: model.Moon,
		LaunchDate:    "2020-01-01",
	}

	service := new(BookingServiceMock)

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, model.ErrBookingValidation)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *EndpointSuite) TestHandleErrorLaunchDate() {
	req := request.BookingRequest{
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "2020-01-01",
		LaunchpadID:   string(model.VandenbergSpaceForceBase1),
		DestinationID: model.Moon,
		LaunchDate:    "xxxx",
	}

	service := new(BookingServiceMock)

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, model.ErrDayDateValidation)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *EndpointSuite) TestHandleErrorBirthday() {
	req := request.BookingRequest{
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "xxxx",
		LaunchpadID:   string(model.VandenbergSpaceForceBase1),
		DestinationID: model.Moon,
		LaunchDate:    "2020-01-01",
	}

	service := new(BookingServiceMock)

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, model.ErrDayDateValidation)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *EndpointSuite) TestHandleSuccess() {
	req := request.BookingRequest{
		Firstname:     "Marcin",
		Lastname:      "",
		Gender:        "",
		Birthday:      "2000-01-01",
		LaunchpadID:   string(model.VandenbergSpaceForceBase1),
		DestinationID: string(model.Moon),
		LaunchDate:    "2020-01-01",
	}

	service := new(BookingServiceMock)
	service.On("Create", mock.Anything).
		Return(nil)
	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)
	s.NoError(err)
}

type BookingServiceMock struct {
	mock.Mock
}

func (b *BookingServiceMock) Create(booking *model.Booking) error {
	args := b.Called(booking)
	return args.Error(0)
}
