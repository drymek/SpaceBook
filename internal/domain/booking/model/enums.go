package model

type LaunchpadID string

const (
	VandenbergSpaceForceBase1 LaunchpadID = "5e9e4501f5090910d4566f83"
	CapeCanaveral1                        = "5e9e4501f509094ba4566f84"
	BocaChicaVillage                      = "5e9e4502f5090927f8566f85"
	OmelekIsland                          = "5e9e4502f5090995de566f86"
	VandenbergSpaceForceBase2             = "5e9e4502f509092b78566f87"
	CapeCanaveral2                        = "5e9e4502f509094188566f88"
)

type DestinationID string

const (
	Mars         DestinationID = "Mars"
	Moon                       = "Moon"
	Pluto                      = "Pluto"
	AsteroidBelt               = "Asteroid Belt"
	Europa                     = "Europa"
	Titan                      = "Titan"
	Ganymede                   = "Ganymede"
)
