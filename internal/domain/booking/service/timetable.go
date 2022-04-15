package service

import (
	"time"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type Timetable map[time.Weekday]model.DestinationID

func NewTimetable() *Timetable {
	return &Timetable{
		time.Monday:    model.Mars,
		time.Tuesday:   model.Moon,
		time.Wednesday: model.Pluto,
		time.Thursday:  model.AsteroidBelt,
		time.Friday:    model.Europa,
		time.Saturday:  model.Titan,
		time.Sunday:    model.Ganymede,
	}
}
