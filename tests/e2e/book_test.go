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
	"dryka.pl/SpaceBook/internal/domain/booking/spacex"
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

	r := repository.NewBookingRepository()
	s.AppDependencies = server.Dependencies{
		Logger:         logger.NewNullLogger(),
		BookingService: service.NewBookingService(r, spacex.NewStaticSpaceXClient()),
		Config:         c,
		Repository:     r,
	}

	s.ValidLaunchpadID = model.VandenbergSpaceForceBase1
	s.InvalidLaunchpadID = "fake-id"
	s.ValidDestinationID = model.AsteroidBelt
	s.InvalidDestinationID = "fake-id"

	suite.Run(t, s)
}

func (s *BookSuite) TestBooking() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	body := fmt.Sprintf(`{
		"id": "123",
		"firstname": "John",
		"lastname": "Doe",
		"gender": "Male",
		"birthday": "2000-07-21",
		"launchpadID": "%s",
		"destinationID": "%s", 
		"launchDate": "2222-01-17"
	}`, s.ValidLaunchpadID, s.ValidDestinationID)

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/bookings", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Contains(string(got), "id")
	s.NotContains(string(got), `{"id":""}`)

	s.Equal(http.StatusOK, res.StatusCode, "Expected status code 200", string(got))
}

func (s *BookSuite) SetupTest() {
	r := repository.NewBookingRepository()
	s.AppDependencies.Repository = r
	s.AppDependencies.BookingService = service.NewBookingService(r, spacex.NewStaticSpaceXClient())
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
		"launchDate": "2222-01-17"
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
		"launchDate": "2222-01-17"
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

func (s *BookSuite) TestBookingList() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()
	date, err := model.NewDayDateFromString("2222-01-17")
	s.NoError(err)
	err = s.AppDependencies.Repository.Create(&model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   "",
		DestinationID: "",
		LaunchDate:    date,
	})
	s.NoError(err)

	res, err := http.Get(srv.URL + "/bookings")
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(`{"items":[{"id":"123","firstname":"Marcin","lastname":"Dryka","gender":"Male","birthday":"2222-01-17","launchpadID":"","destinationID":"","launchDate":"2222-01-17"}]}
`, string(got))
	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode, "Expected status code 200")
}

func (s *BookSuite) TestBookingDelete() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()
	date, err := model.NewDayDateFromString("2222-01-17")
	s.NoError(err)
	booking := &model.Booking{
		ID:            "123",
		Firstname:     "Marcin",
		Lastname:      "Dryka",
		Gender:        "Male",
		Birthday:      date,
		LaunchpadID:   "",
		DestinationID: "",
		LaunchDate:    date,
	}
	err = s.AppDependencies.Repository.Create(booking)
	s.NoError(err)

	req, err := http.NewRequest(http.MethodDelete, srv.URL+"/bookings/"+booking.ID, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")
	s.NoError(err)
	res, err := http.DefaultClient.Do(req)
	s.NoError(err)

	got, err := ioutil.ReadAll(res.Body)
	_ = got
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(http.StatusNoContent, res.StatusCode, "Expected status code 204")
}
