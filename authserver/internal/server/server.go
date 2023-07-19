package server

import (
	"authserver/pkg/logging"
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

// Start - запуск сервера
func (s *Server) Start() error {
	//if len(s.cfg.CryptoKey) > 0 {
	//	return s.httpServer.ListenAndServeTLS(certFile, s.cfg.CryptoKey)
	//}
	return s.httpServer.ListenAndServe()
}

// GetServer - получение сервера
func (s *Server) GetServer() *http.Server {
	return s.httpServer
}
