package server

import (
	"dryka.pl/SpaceBook/internal/application/config"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
)

type Dependencies struct {
	Logger         logger.Logger
	BookingService service.BookingService
	Config         config.Config
}
