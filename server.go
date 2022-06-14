package todo

import (
	"errors"
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/handler"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(c *handler.Controller) (*Server, error) {
	if len(c.Router.Routes) == 0 {
		return nil, errors.New("apply at least one route")
	}

	s := Server{}
	s.httpServer = &http.Server{
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 5,
		WriteTimeout:   time.Second * 5,
		Handler:        c.Router.Handler,
	}

	return &s, nil
}

func (s *Server) StartListeningOn(port int) error {

	s.httpServer.Addr = fmt.Sprintf(":%d", port)
	fmt.Printf("server has successfully runnning on port %d \n", port)
	return s.httpServer.ListenAndServe()
}
