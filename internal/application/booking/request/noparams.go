package request

import (
	"context"
	"net/http"

	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	httpkit "github.com/go-kit/kit/transport/http"
)

func DecodeListRequest(_ logger.Logger) httpkit.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		return nil, nil
	}
}
