package server

import (
	"GophKeeper/keepserver/pkg/logging"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	logger     logging.Logger
}

func NewServer(log logging.Logger, address string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    address,
			Handler: handler,
		},
		logger: log,
	}
}
