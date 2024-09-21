package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Start() (srvCh chan error)
	Stop() error
}

type server struct {
	srv           *http.Server
	srvCh         chan error
	stopTimeoutMS time.Duration
}

func (s *server) Start() (srvCh chan error) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.srvCh <- err
		}
	}()

	return s.srvCh
}

func (s *server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.stopTimeoutMS)

	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %v", err)
	}
	return nil
}

func New(gin *Router) (Server, error) {
	s := &server{
		srv: &http.Server{
			Addr:    ":8080",
			Handler: gin.Engine,
		},
		stopTimeoutMS: 5000 * time.Millisecond,
	}

	return s, nil
}
