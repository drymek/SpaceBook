package service

import (
	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type SpaceXClient interface {
	GetLaunches(date model.DayDate, id model.LaunchpadID) ([]string, error)
}

type staticSpaceXClient struct {
}

func (s staticSpaceXClient) GetLaunches(date model.DayDate, id model.LaunchpadID) ([]string, error) {
	if date.Format("2006-01-02") == "2222-02-22" {
		return []string{"Falcon 9"}, nil
	}
	return []string{}, nil
}

func NewStaticSpaceXClient() SpaceXClient {
	return &staticSpaceXClient{}
}
