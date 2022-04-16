package endpoint

import (
	"context"
	"net/http"

	requestx "dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/application/booking/response"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"github.com/go-kit/kit/endpoint"
)

func MakeCreateEndpoint(_ logger.Logger, service service.BookingService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		r := request.(requestx.BookingRequest)

		birthday, err := model.NewDayDateFromString(r.Birthday)
		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		launchDate, err := model.NewDayDateFromString(r.LaunchDate)
		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		booking, err := model.NewBooking(
			r.ID,
			r.Firstname,
			r.Lastname,
			r.Gender,
			birthday,
			model.LaunchpadID(r.LaunchpadID),
			model.DestinationID(r.DestinationID),
			launchDate,
		)
		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		err = service.Create(&booking)
		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		return response.NewBookingIdResponse(http.StatusCreated, booking), nil
	}
}
