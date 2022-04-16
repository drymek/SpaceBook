package request

import (
	"context"
	"encoding/json"
	"net/http"

	"dryka.pl/SpaceBook/internal/application/httpx"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	httpkit "github.com/go-kit/kit/transport/http"
)

type BookingRequest struct {
	ID            string `json:"id"`
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Gender        string `json:"gender"`
	Birthday      string `json:"birthday"`
	LaunchpadID   string `json:"launchpadID"`
	DestinationID string `json:"destinationID"`
	LaunchDate    string `json:"launchDate"`
}

func DecodeBookingRequest(logger logger.Logger) httpkit.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var request BookingRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			if err2 := logger.Log("function", "DecodeBookingRequest", "error", err); err2 != nil {
				return nil, err2
			}

			return nil, httpx.NewBadRequest(err)
		}

		return request, nil
	}
}
