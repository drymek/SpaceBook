package e2e

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"dryka.pl/SpaceBook/internal/application/config"
	"dryka.pl/SpaceBook/internal/application/server"
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/service"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"dryka.pl/SpaceBook/internal/infrastructure/persistence/inmemory/repository"
	"github.com/stretchr/testify/suite"
)

type BookSuite struct {
	suite.Suite
	AppDependencies      server.Dependencies
	ValidLaunchpadID     model.LaunchpadID
	InvalidLaunchpadID   model.LaunchpadID
	InvalidDestinationID model.DestinationID
	ValidDestinationID   model.DestinationID
}

func TestBookSuite(t *testing.T) {
	s := new(BookSuite)
	c, err := config.NewConfig()
	if err != nil {
		t.Fatalf("invalid config: %v", err)
	}

	s.AppDependencies = server.Dependencies{
		Logger:         logger.NewNullLogger(),
		BookingService: service.NewBookingService(repository.NewBookingRepository(), service.NewStaticSpaceXClient()),
		Config:         c,
	}

	s.ValidLaunchpadID = model.VandenbergSpaceForceBase1
	s.InvalidLaunchpadID = "fake-id"
	s.ValidDestinationID = model.Mars
	s.InvalidDestinationID = "fake-id"

	suite.Run(t, s)
}

func (s *BookSuite) TestBooking() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	body := fmt.Sprintf(`{
		"firstname": "John",
		"lastname": "Doe",
		"gender": "Male",
		"birthday": "2000-07-21",
		"launchpadID": "%s",
		"destinationID": "%s", 
		"launchDate": "2022-01-17"
	}`, s.ValidLaunchpadID, s.ValidDestinationID)

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/bookings", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Contains(string(got), "id")

	s.Equal(http.StatusOK, res.StatusCode, "Expected status code 200", string(got))
}

func (s *BookSuite) TestInvalidBookingLaunchpad() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	body := fmt.Sprintf(`{
		"firstname": "John",
		"lastname": "Doe",
		"gender": "Male",
		"birthday": "2000-07-21",
		"launchpadID": "%s",
		"destinationID": "%s", 
		"launchDate": "2022-01-20"
	}`, s.InvalidLaunchpadID, s.ValidDestinationID)

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/bookings", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(`{"error":"booking validation error: invalid launchpad_id"}`, string(got))
	s.Nil(err)
	s.Equal(http.StatusBadRequest, res.StatusCode, "Expected status code 400")
}

func (s *BookSuite) TestInvalidBookingDestination() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	body := fmt.Sprintf(`{
		"firstname": "John",
		"lastname": "Doe",
		"gender": "Male",
		"birthday": "2000-07-21",
		"launchpadID": "%s",
		"destinationID": "%s", 
		"launchDate": "2022-01-20"
	}`, s.ValidLaunchpadID, s.InvalidDestinationID)

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/bookings", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(`{"error":"booking validation error: invalid destination_id"}`, string(got))
	s.Nil(err)
	s.Equal(http.StatusBadRequest, res.StatusCode, "Expected status code 400")
}
