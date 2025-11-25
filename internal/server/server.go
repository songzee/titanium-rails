package server

import (
	"context"
	"fmt"
	"net/http"
	"titanium-rails/internal/config"
	"titanium-rails/internal/handlers"
	"titanium-rails/internal/ledger"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
	http   *http.Server
	ledger ledger.Ledger
}

func New(cfg *config.Config) *Server {
	r := chi.NewRouter()
	
	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	// Create ledger stub
	l := ledger.NewStub()

	// Create handlers
	h := handlers.New(l)

	// Routes
	r.Get("/health", handlers.Health)
	r.Post("/stms/auth", h.STMSAuth)

	s := &Server{
		router: r,
		ledger: l,
	}

	s.http = &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	return s
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

