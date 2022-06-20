package server

import (
	"errors"
	"fmt"
	http2 "github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(h *http2.Handler) (*Server, error) {
	if len(h.Router.Routes) == 0 {
		return nil, errors.New("apply at least one route")
	}

	s := Server{}
	s.httpServer = &http.Server{
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 5,
		WriteTimeout:   time.Second * 5,
		Handler:        h.Router.Handler,
	}

	return &s, nil
}

func (s *Server) StartListeningOn(port int) error {
	s.httpServer.Addr = fmt.Sprintf(":%d", port)
	return s.httpServer.ListenAndServe()
}
