package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/health", app.healthHandler)
	router.Post("/api/users", app.handlers.User.CreateUser)

	return router
}

func (app *application) run(handler http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: handler,
	}

	fmt.Printf("Starting server on %s\n", app.config.addr)
	return srv.ListenAndServe()
}
