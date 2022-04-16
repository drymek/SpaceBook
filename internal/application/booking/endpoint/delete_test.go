package endpoint_test

import (
	"context"
	"net/http"
	"testing"

	"dryka.pl/SpaceBook/internal/application/booking/endpoint"
	"dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
	service2 "dryka.pl/SpaceBook/internal/domain/booking/service"
	"github.com/stretchr/testify/suite"
)

type DeleteEndpointSuite struct {
	suite.Suite
}

func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteEndpointSuite))
}

func (s *DeleteEndpointSuite) TestHandleNotFoundErrorFromService() {
	req := request.BookingIDRequest{
		ID: "123",
	}

	serviceErr := service2.ErrBookingService(repository.ErrBookingNotFound)
	service := new(BookingServiceMock)
	service.On("Delete", "123").
		Return(serviceErr)

	_, err := endpoint.MakeDeleteEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, serviceErr)
	s.Equal(http.StatusNotFound, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *DeleteEndpointSuite) TestHandleErrorFromService() {
	req := request.BookingIDRequest{
		ID: "123",
	}

	serviceErr := service2.ErrBookingService(repository.ErrBookingAlreadyExists)
	service := new(BookingServiceMock)
	service.On("Delete", "123").
		Return(serviceErr)

	_, err := endpoint.MakeDeleteEndpoint(nil, service)(context.TODO(), req)

	s.Error(err)
	s.ErrorIs(err, serviceErr)
	s.Equal(http.StatusInternalServerError, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *DeleteEndpointSuite) TestHandleSuccess() {
	req := request.BookingIDRequest{
		ID: "123",
	}

	service := new(BookingServiceMock)
	service.On("Delete", "123").
		Return(nil)
	_, err := endpoint.MakeDeleteEndpoint(nil, service)(context.TODO(), req)
	s.NoError(err)
}
