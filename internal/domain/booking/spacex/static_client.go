package spacex

import (
	"context"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type staticSpaceXClient struct {
}

func (s staticSpaceXClient) GetLaunches(ctx context.Context, date model.DayDate, id model.LaunchpadID) ([]string, error) {
	if date.Format("2006-01-02") == "2222-02-22" {
		return []string{"Falcon 9"}, nil
	}
	return []string{}, nil
}

func NewStaticSpaceXClient() SpaceXClient {
	return &staticSpaceXClient{}
}
