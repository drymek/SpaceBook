package service

import (
	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type SpaceXClient interface {
	GetLaunches(date model.DayDate, id model.LaunchpadID) ([]string, error)
}
