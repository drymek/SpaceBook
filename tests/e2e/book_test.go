package e2e

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"dryka.pl/SpaceBook/internal/application/server"
	"github.com/stretchr/testify/suite"
)

type BookSuite struct {
	suite.Suite
	AppDependencies server.Dependencies
}

func TestBookSuite(t *testing.T) {
	s := new(BookSuite)
	s.AppDependencies = server.Dependencies{}

	suite.Run(t, s)
}

func (s *BookSuite) TestBooking() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	//@TODO
	body := `{
		firstname: "John",
		lastname: "Doe",
		gender: "Male",
		birthday: "2000-07-21",
		launchpadID: "FREE LAUNCHPAD ID",
		destinationID: "FREE LAUNCHPAD ID", 
		launchDate: ""
	}`

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/book", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Contains(string(got), "id")

	s.Nil(err)
	s.Equal(http.StatusNoContent, res.StatusCode, "Expected status code 204")
}

func (s *BookSuite) TestInvalidBookingLaunchpad() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	//@TODO
	body := `{
		firstname: "John",
		lastname: "Doe",
		gender: "Male",
		birthday: "2000-07-21",
		launchpadID: "TAKEN LAUNCHPAD ID",
		destinationID: "FREE LAUNCHPAD ID", 
		launchDate: ""
	}`

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/book", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(`{"error":"Launchpad is already booked"}`, string(got))
	s.Nil(err)
	s.Equal(http.StatusBadRequest, res.StatusCode, "Expected status code 400")
}

func (s *BookSuite) TestInvalidBookingDestination() {
	srv := httptest.NewServer(server.NewServer(s.AppDependencies))
	defer srv.Close()

	//@TODO
	body := `{
		firstname: "John",
		lastname: "Doe",
		gender: "Male",
		birthday: "2000-07-21",
		launchpadID: "TAKEN LAUNCHPAD ID",
		destinationID: "FREE LAUNCHPAD ID", 
		launchDate: ""
	}`

	requestBody := []byte(body)
	res, err := http.Post(srv.URL+"/book", "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	got, err := ioutil.ReadAll(res.Body)
	err2 := res.Body.Close()
	s.Nil(err)
	s.Nil(err2)

	s.Equal(`{"error":"Invalid destination"}`, string(got))
	s.Nil(err)
	s.Equal(http.StatusBadRequest, res.StatusCode, "Expected status code 400")
}
