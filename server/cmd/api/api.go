package main

import (
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (a *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", a.healthcheckHandler)
	return mux
}

func (a *application) run() error {

	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      a.mount(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()
}
