package endpoint

import (
	"context"
	"errors"
	"net/http"

	requestx "dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/application/booking/response"
	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"github.com/go-kit/kit/endpoint"
)

func MakeDeleteEndpoint(_ logger.Logger, service service.BookingService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		r := request.(requestx.BookingIDRequest)

		err := service.Delete(r.ID)

		if err != nil {
			if errors.Is(err, repository.ErrBookingNotFound) {
				return nil, httpx.NewNotFound(err)
			}

			return nil, httpx.NewInternalServerError(err)
		}

		return response.NewEmptyResponse(http.StatusNoContent), nil
	}
}
