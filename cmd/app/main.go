package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"dryka.pl/SpaceBook/internal/application/config"
	"dryka.pl/SpaceBook/internal/application/server"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"dryka.pl/SpaceBook/internal/infrastructure/persistence/inmemory/repository"
)

func main() {
	l := logger.NewLogger()
	err := l.Log("msg", "Starting service")
	if err != nil {
		panic(err)
	}

	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	r := repository.NewBookingRepository()
	dependencies := server.Dependencies{
		Logger:         l,
		BookingService: service.NewBookingService(r, service.NewStaticSpaceXClient()),
		Config:         c,
		Repository:     r,
	}

	muxer := http.TimeoutHandler(server.NewServer(dependencies), c.GetTimeout(), "Timeout!")
	srv := http.Server{Addr: c.GetHttpAddr(), Handler: muxer}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		if err := srv.Shutdown(ctx); err != nil {
			err := l.Log("HTTP server Shutdown: %v", err)
			if err != nil {
				panic(err)
			}
		}
	}()

	err = l.Log("transport", "http", "address", c.GetHttpAddr(), "msg", "listening")
	if err != nil {
		panic(err)
	}

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		err := l.Log("HTTP server ListenAndServe: %v", err)
		if err != nil {
			panic(err)
		}
	}
}
