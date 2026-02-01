package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/saianand32/Quipify/internal/config"
)

type application struct {
	config *config.Config
}

func (a *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Logger)

	r.Route("/v1", func(e chi.Router) {
		r.Get("/health", a.healthcheckHandler)
	})
	return r
}

func (a *application) run() error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.config.Port),
		Handler:      a.mount(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
