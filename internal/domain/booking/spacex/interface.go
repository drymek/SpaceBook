package spacex

import (
	"context"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type SpaceXClient interface {
	GetLaunches(ctx context.Context, date model.DayDate, id model.LaunchpadID) ([]string, error)
}
