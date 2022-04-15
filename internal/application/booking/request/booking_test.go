package request_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/infrastructure/logger"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
)

type BookingSuite struct {
	suite.Suite
}

func TestRequestSuite(t *testing.T) {
	suite.Run(t, new(BookingSuite))
}

func (s *BookingSuite) TestValidInput() {
	json := `{
		"firstname": "John",
		"lastname": "Doe",
		"gender": "Male",
		"birthday": "2000-07-21",
		"launchpadID": "FREE LAUNCHPAD ID",
		"destinationID": "FREE LAUNCHPAD ID", 
		"launchDate": "2022-07-22"
	}`

	req := httptest.NewRequest(http.MethodPost, "/booking", strings.NewReader(json))
	got, err := request.DecodeBookingRequest(logger.NewNullLogger())(context.Background(), req)

	s.Nil(err)

	want := request.BookingRequest{
		Firstname:     "John",
		Lastname:      "Doe",
		Gender:        "Male",
		Birthday:      "2000-07-21",
		LaunchpadID:   "FREE LAUNCHPAD ID",
		DestinationID: "FREE LAUNCHPAD ID",
		LaunchDate:    "2022-07-22",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		s.T().Fatalf("mismatch (-want +got):\n%v", diff)
	}
}

func (s *BookingSuite) TestInvalidInput() {
	json := `
		gender: "Male",
		birthday: "2000-07-21",
	`

	req := httptest.NewRequest(http.MethodPost, "/booking", strings.NewReader(json))
	_, err := request.DecodeBookingRequest(logger.NewNullLogger())(context.Background(), req)

	s.Error(err)
}
