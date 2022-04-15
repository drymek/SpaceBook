package server

import (
	"net/http"

	"dryka.pl/SpaceBook/internal/application/booking/endpoint"
	"dryka.pl/SpaceBook/internal/application/booking/request"
	"dryka.pl/SpaceBook/internal/application/httpx"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewServer(d Dependencies) *mux.Router {
	r := mux.NewRouter()

	createHandler := kithttp.NewServer(
		endpoint.MakeCreateEndpoint(d.Logger, d.BookingService),
		request.DecodeBookingRequest(d.Logger),
		httpx.EncodeResponse(d.Logger),
		kithttp.ServerErrorEncoder(httpx.EncodeError(d.Logger)),
	)

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(d.Config.GetNewRelicConfigAppName()),
		newrelic.ConfigLicense(d.Config.GetNewRelicConfigLicense()),
	)

	if err != nil {
		err2 := d.Logger.Log("context", "newrelic", "error", err)
		if err2 != nil {
			panic(err2)
		}
	}

	r.Handle(newrelic.WrapHandle(app, "/bookings", AccessControl(createHandler))).Methods(http.MethodOptions, http.MethodPost)

	return r
}
