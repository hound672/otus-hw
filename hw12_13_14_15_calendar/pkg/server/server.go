package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/utrack/clay/v3/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hound672/otus-hw/hw12_13_14_15_calendar/pkg/logger"
)

type Server struct {
	wg sync.WaitGroup

	grpcListener net.Listener
	grpcServer   *grpc.Server

	httpListener net.Listener
	httpRouter   chi.Router
	httpServer   *http.Server
}

const (
	shutdownServerTimeout = 10 * time.Second
)

func New(config *Config) (*Server, error) {
	// building grpc server
	grpcServer := grpc.NewServer()

	if config.UseReflection {
		reflection.Register(grpcServer)
	}

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PortGRPC))
	if err != nil {
		return nil, fmt.Errorf("net.Listen (grpc): %w", err)
	}
	httpListener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.PortHTTP))
	if err != nil {
		return nil, fmt.Errorf("net.Listen (http): %w", err)
	}

	// building http server
	hmux := chi.NewRouter()
	httpServer := &http.Server{
		Handler:           hmux,
		ReadTimeout:       time.Duration(config.HTTPReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(config.HTTPWriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(config.HTTPIdleTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(config.HTTPReadHeaderTimeout) * time.Second,
	}

	s := &Server{
		grpcServer:   grpcServer,
		grpcListener: grpcListener,

		httpListener: httpListener,
		httpRouter:   hmux,
		httpServer:   httpServer,
	}
	return s, nil
}

func (s *Server) Run(controllers ...transport.Service) error {
	for _, controller := range controllers {
		controller := controller

		controller.GetDescription().RegisterGRPC(s.grpcServer)
		controller.GetDescription().RegisterHTTP(s.httpRouter)
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		logger.Info("Start GRPC server", "addr", s.grpcListener.Addr())
		if err := s.grpcServer.Serve(s.grpcListener); err != nil {
			logger.Error("s.grpcServer.Serve", "err", err)
			return
		}
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		logger.Info("Start HTTP server", "addr", s.httpListener.Addr())
		if err := s.httpServer.Serve(s.httpListener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("s.httpServer.Serve", "err", err)
			return
		}
	}()

	return nil
}

func (s *Server) Stop() {
	s.grpcServer.Stop()

	ctx, cancel := context.WithTimeout(context.Background(), shutdownServerTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Error("s.httpServer.Shutdown", "err", err)
	}

	s.wg.Wait()
}
