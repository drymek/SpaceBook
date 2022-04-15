package model

type LaunchpadID string

const (
	VandenbergSpaceForceBase1 LaunchpadID = "5e9e4501f5090910d4566f83"
	CapeCanaveral1                        = "5e9e4501f509094ba4566f84"
	BocaChicaVillage                      = "5e9e4501f509094ba4566f85"
	OmelekIsland                          = "5e9e4501f509094ba4566f86"
	VandenbergSpaceForceBase2             = "5e9e4501f509094ba4566f87"
	CapeCanaveral2                        = "5e9e4501f509094ba4566f88"
)

type DestinationID string

const (
	Mars     DestinationID = "Mars"
	Moon                   = "Moon"
	Pluto                  = "Pluto"
	Asteroid               = "Asteroid"
	Belt                   = "Belt"
	Europa                 = "Europa"
	Titan                  = "Titan"
	Ganymede               = "Ganymede"
)
