package server

import "net/http"

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}
