package spacex

import (
	"context"
	"testing"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
	"github.com/stretchr/testify/suite"
)

type SpacexClientSuite struct {
	suite.Suite
}

func TestDayDateSuite(t *testing.T) {
	s := new(SpacexClientSuite)

	suite.Run(t, s)
}

func (s *SpacexClientSuite) TestGetEmptyLaunches() {
	ctx := context.Background()
	date, err := model.NewDayDateFromString("2020-06-01")
	s.NoError(err)
	launches, err := NewSpaceXClient().GetLaunches(ctx, date, model.VandenbergSpaceForceBase1)
	s.NoError(err)
	s.Empty(launches)
}

func (s *SpacexClientSuite) TestGetLaunches() {
	ctx := context.Background()
	date, err := model.NewDayDateFromString("2014-04-18")
	s.NoError(err)
	launches, err := NewSpaceXClient().GetLaunches(ctx, date, model.CapeCanaveral1)
	s.NoError(err)
	s.NotEmpty(launches)
}

func (s *SpacexClientSuite) TestGetLaunchesError() {
	ctx := context.Background()
	date, err := model.NewDayDateFromString("2020-06-01")
	s.NoError(err)

	client := &spaceXClient{
		baseUrl:               "https://example.com",
		queryLaunchesEndpoint: "/v4/launches/query",
	}
	_, err = client.GetLaunches(ctx, date, model.VandenbergSpaceForceBase1)
	s.Error(err)
}
