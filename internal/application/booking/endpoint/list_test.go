package endpoint_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"dryka.pl/SpaceBook/internal/application/booking/endpoint"
	"dryka.pl/SpaceBook/internal/application/booking/response"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	service2 "dryka.pl/SpaceBook/internal/domain/booking/service"
	"github.com/stretchr/testify/suite"
)

type ListEndpointSuite struct {
	suite.Suite
}

func TestListSuite(t *testing.T) {
	suite.Run(t, new(ListEndpointSuite))
}

func (s *ListEndpointSuite) TestHandleErrorFromService() {
	serviceErr := service2.ErrBookingService(fmt.Errorf("test error"))
	service := new(BookingServiceMock)
	service.On("List").
		Return([]model.Booking{}, serviceErr)

	_, err := endpoint.MakeListEndpoint(nil, service)(context.TODO(), nil)

	s.Error(err)
	s.ErrorIs(err, serviceErr)
	s.Equal(http.StatusInternalServerError, err.(httpx.StatusCodeHolder).StatusCode())
}

func (s *ListEndpointSuite) TestSuccessfulList() {
	service := new(BookingServiceMock)
	date, err := model.NewDayDateFromString("2022-02-22")
	s.NoError(err)
	models := []model.Booking{
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
	service.On("List").
		Return(models, nil)

	bookings, err := endpoint.MakeListEndpoint(nil, service)(context.TODO(), nil)

	s.NoError(err)
	s.Equal(http.StatusOK, bookings.(httpx.StatusCodeHolder).StatusCode())
	s.Equal(response.NewBookingResponse(http.StatusOK, models), bookings)
}
