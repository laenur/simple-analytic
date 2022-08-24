package server

import (
	"fmt"
	"net/http"

	"github.com/laenur/simple-analytic/pkg/logger"
)

type server struct {
	port int
	mux  *http.ServeMux
}

func NewServer(port int) server {
	return server{
		port: port,
		mux:  http.NewServeMux(),
	}
}

func (s *server) AddHandler(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func (s server) StartServer() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.mux,
	}

	logger.Info("Started", "Port", s.port)
	return server.ListenAndServe()
}
