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
		Birthday:      "",
		LaunchpadID:   "",
		DestinationID: "",
		LaunchDate:    "",
	}

	service := new(BookingServiceMock)
	service.On("Create", mock.Anything).
		Return(service2.ErrInvalidBookingDate(nil))

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *EndpointSuite) TestHandleErrorFromBooking() {
	req := request.BookingRequest{}

	service := new(BookingServiceMock)

	_, err := endpoint.MakeCreateEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.Equal(http.StatusBadRequest, err.(httpx.StatusCodeHolder).StatusCode())
}

type BookingServiceMock struct {
	mock.Mock
}

func (b *BookingServiceMock) Create(booking model.Booking) error {
	args := b.Called(booking)
	return args.Error(0)
}
