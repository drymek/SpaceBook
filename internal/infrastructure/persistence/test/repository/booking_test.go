package repository

import (
	"database/sql"
	"sync"
	"testing"

	"dryka.pl/SpaceBook/internal/application/config"
	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"dryka.pl/SpaceBook/internal/domain/booking/repository"
	inmemory "dryka.pl/SpaceBook/internal/infrastructure/persistence/inmemory/repository"
	"dryka.pl/SpaceBook/internal/infrastructure/persistence/postgres"
	postgresx "dryka.pl/SpaceBook/internal/infrastructure/persistence/postgres/repository"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
)

type RepositoryIntegrationSuite struct {
	suite.Suite
	repositories map[string]repository.BookingRepository
	mu           sync.Mutex
	client       *sql.DB
}

func TestServiceSuite(t *testing.T) {
	s := new(RepositoryIntegrationSuite)
	c, err := config.NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	s.client, err = postgres.NewConnection(
		c.GetDatabase().GetDatabaseHost(),
		c.GetDatabase().GetDatabasePort(),
		c.GetDatabase().GetDatabaseUser(),
		c.GetDatabase().GetDatabasePassword(),
		c.GetDatabase().GetDatabaseName(),
	)
	if err != nil {
		t.Fatal(err)
	}

	s.repositories = make(map[string]repository.BookingRepository)
	s.repositories["postgres"] = postgresx.NewBookingRepository(s.client)
	s.repositories["inmemory"] = inmemory.NewBookingRepository()
	s.mu = sync.Mutex{}
	suite.Run(t, s)
}

func (s *RepositoryIntegrationSuite) SetupTest() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.repositories["inmemory"] = inmemory.NewBookingRepository()
	_, err := s.client.Exec("DELETE FROM bookings")
	s.NoError(err)
}

func (s *RepositoryIntegrationSuite) TestCreateAndFetch() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	for i := range s.repositories {
		s.T().Run(i, func(t *testing.T) {
			s.mu.Lock()
			defer s.mu.Unlock()
			booking := &model.Booking{
				ID:            "123",
				Firstname:     "Marcin",
				Lastname:      "Dryka",
				Gender:        "Male",
				Birthday:      date,
				LaunchpadID:   model.VandenbergSpaceForceBase1,
				DestinationID: model.Moon,
				LaunchDate:    date,
			}

			err := s.repositories[i].Create(booking)
			s.NoError(err)
			got, err := s.repositories[i].Find(booking.ID)
			s.NoError(err)
			want := *booking
			if diff := cmp.Diff(want, got); diff != "" {
				s.Failf("value mismatch", "(-want +got):\n%v", diff)
			}
		})
	}
}

func (s *RepositoryIntegrationSuite) TestList() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	for i := range s.repositories {
		s.T().Run(i, func(t *testing.T) {
			s.mu.Lock()
			defer s.mu.Unlock()
			booking := &model.Booking{
				ID:            "123",
				Firstname:     "Marcin",
				Lastname:      "Dryka",
				Gender:        "Male",
				Birthday:      date,
				LaunchpadID:   model.VandenbergSpaceForceBase1,
				DestinationID: model.Moon,
				LaunchDate:    date,
			}
			got, err := s.repositories[i].List()
			s.Empty(got)
			s.NoError(err)

			err = s.repositories[i].Create(booking)

			s.NoError(err)
			got, err = s.repositories[i].List()
			s.NoError(err)
			s.NotEmpty(got)
			s.Len(got, 1)
		})
	}
}

func (s *RepositoryIntegrationSuite) TestDelete() {
	date, err := model.NewDayDateFromString("2020-01-01")
	s.NoError(err)
	for i := range s.repositories {
		s.T().Run(i, func(t *testing.T) {
			s.mu.Lock()
			defer s.mu.Unlock()
			booking := &model.Booking{
				ID:            "123",
				Firstname:     "Marcin",
				Lastname:      "Dryka",
				Gender:        "Male",
				Birthday:      date,
				LaunchpadID:   model.VandenbergSpaceForceBase1,
				DestinationID: model.Moon,
				LaunchDate:    date,
			}

			err := s.repositories[i].Create(booking)
			s.NoError(err)

			err = s.repositories[i].Delete(booking.ID)
			s.NoError(err)

			_, err = s.repositories[i].Find(booking.ID)
			s.Error(err)
			s.Equal(repository.ErrBookingNotFound, err)
		})
	}
}
