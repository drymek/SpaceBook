package spacex

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"dryka.pl/SpaceBook/internal/domain/booking/model"
)

type spaceXClient struct {
	baseUrl               string
	queryLaunchesEndpoint string
	defaultTimeout        time.Duration
}

type launchpad struct {
	ID string `json:"launchpad"`
}

type result struct {
	Docs []launchpad
}

func (r result) getLaunches() []string {
	var launches []string
	for _, doc := range r.Docs {
		launches = append(launches, doc.ID)
	}
	return launches
}

func (s spaceXClient) GetLaunches(ctx context.Context, date model.DayDate, id model.LaunchpadID) ([]string, error) {
	query := s.getLaunchesQuery(id, date)

	got, err := s.doRequest(ctx, http.MethodPost, s.queryLaunchesEndpoint, query)
	if err != nil {
		return []string{}, fmt.Errorf("spacex client error: %w", err)
	}

	var res result
	if err := json.Unmarshal(got, &res); err != nil {
		return []string{}, fmt.Errorf("spacex client error: %w", err)
	}
	return res.getLaunches(), nil
}

func (s spaceXClient) doRequest(ctx context.Context, method string, endpoint string, query string) ([]byte, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		fmt.Sprintf("%s%s", s.baseUrl, endpoint),
		strings.NewReader(query),
	)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: s.defaultTimeout}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return got, nil
}

func (s spaceXClient) getLaunchesQuery(id model.LaunchpadID, date model.DayDate) string {
	layout := "2006-01-02T15:04:05.000Z"
	query := fmt.Sprintf(`
		{
		  "query": {
			"launchpad": {
			  "$eq": "%s"
			},
			"date_utc": {
			  "$gte": "%s",
			  "$lte": "%s"
			}
		  },
		  "options": {}
		}
	`, id, date.Start().Format(layout), date.End().Format(layout))

	return query
}

func NewSpaceXClient() SpaceXClient {
	return &spaceXClient{
		baseUrl:               "https://api.spacexdata.com",
		queryLaunchesEndpoint: "/v4/launches/query",
		defaultTimeout:        time.Second * 10,
	}
}
