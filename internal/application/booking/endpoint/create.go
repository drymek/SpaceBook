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

		booking, err := model.NewBooking(
			r.Firstname,
			r.Lastname,
			r.Gender,
			r.Birthday,
			model.LaunchpadID(r.LaunchpadID),
			model.DestinationID(r.DestinationID),
			r.LaunchDate,
		)

		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		err = service.Create(booking)

		if err != nil {
			return nil, httpx.NewBadRequest(err)
		}

		return response.NewBookingIdResponse(http.StatusCreated, booking), nil
	}
}
