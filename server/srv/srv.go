package srv

import (
	"context"
	"log"
	"net/http"
)

type server struct {
	sv http.Server
}

type Server interface {
	Run()
	Stop(ctx context.Context) error
}

func NewServer(port string, handler http.Handler) Server {
	return &server{
		sv: http.Server{
			Addr:    port,
			Handler: handler,
		},
	}
}

func (s *server) Run() {
	go func() {
		if err := s.sv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()
}

func (s *server) Stop(ctx context.Context) error {
	return s.sv.Shutdown(ctx)
}
