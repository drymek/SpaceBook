package healthcheck

import (
	"context"

	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoint(_ logger.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return Response{}, nil
	}
}
