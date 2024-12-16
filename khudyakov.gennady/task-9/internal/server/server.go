package server

import (
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(config *ServerConfig, router http.Handler) *Server {
	httpServer := &http.Server{
		Addr:    config.Host,
		Handler: router,
	}

	serv := &Server{
		server: httpServer,
	}

	return serv
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
