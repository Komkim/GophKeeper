package server

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Server struct {
	httpServer *http.Server
	log        *zerolog.Event
}

func NewServer(log *zerolog.Event, address string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    address,
			Handler: handler,
		},
		log: log,
	}
}
