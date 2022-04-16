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
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type BookingIDRequestSuite struct {
	suite.Suite
}

func TestBookingIDRequestSuite(t *testing.T) {
	suite.Run(t, new(BookingIDRequestSuite))
}

func (s *BookingIDRequestSuite) TestNoIdInput() {
	req := httptest.NewRequest(http.MethodPost, "/bookings", strings.NewReader(`{}`))
	ctx := context.Background()
	values := map[string]string{}
	mux.SetURLVars(req, values)
	_, err := request.DecodeDeleteRequest(logger.NewNullLogger())(ctx, req)

	s.Error(err)
	s.ErrorIs(err, request.ErrInvalidBookingID)
}

func (s *BookingIDRequestSuite) TestValidInput() {
	req := httptest.NewRequest(http.MethodPost, "/bookings/{id}", strings.NewReader(`{}`))
	ctx := context.Background()
	values := map[string]string{
		"id": "123",
	}
	req = mux.SetURLVars(req, values)
	got, err := request.DecodeDeleteRequest(logger.NewNullLogger())(ctx, req)

	s.Nil(err)

	want := request.BookingIDRequest{
		ID: "123",
	}

	if diff := cmp.Diff(want, got); diff != "" {
		s.T().Fatalf("mismatch (-want +got):\n%v", diff)
	}
}
