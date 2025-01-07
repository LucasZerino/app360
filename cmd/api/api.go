package main

import (
	"fmt"
	"net/http"

	"app360/cmd/api/middleware"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
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

	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)

	// Rotas públicas
	router.Post("/api/users", app.handlers.User.CreateUser)
	router.Post("/api/login", app.handlers.Auth.Login)

	// Rotas protegidas
	router.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(app.jwt))

		r.Post("/api/logout", app.handlers.Auth.Logout)
		// Adicione outras rotas protegidas aqui
	})

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
