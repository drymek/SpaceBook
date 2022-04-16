package endpoint

import (
	"context"
	"net/http"

	"dryka.pl/SpaceBook/internal/application/booking/response"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"github.com/go-kit/kit/endpoint"
)

func MakeListEndpoint(_ logger.Logger, service service.BookingService) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		bookings, err := service.List()
		if err != nil {
			return nil, httpx.NewInternalServerError(err)
		}

		return response.NewBookingResponse(http.StatusOK, bookings), nil
	}
}
