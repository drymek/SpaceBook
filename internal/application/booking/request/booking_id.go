package request

import (
	"context"
	"fmt"
	"net/http"

	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	httpkit "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type BookingIDRequest struct {
	ID string
}

var ErrInvalidBookingID = fmt.Errorf("invalid booking id")

func DecodeDeleteRequest(logger logger.Logger) httpkit.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var request BookingIDRequest

		id, ok := mux.Vars(r)["id"]
		if !ok {
			return nil, httpx.NewBadRequest(ErrInvalidBookingID)
		}

		request.ID = id

		return request, nil
	}
}
