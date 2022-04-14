package server

import (
	"github.com/gorilla/mux"
)

func NewServer(d Dependencies) *mux.Router {
	r := mux.NewRouter()

	return r
}
