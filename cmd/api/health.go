package main

import (
	"net/http"
	"time"

	"app360/internal/store"

	"github.com/google/uuid"
)

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))

	app.store.User.Create(r.Context(), &store.User{
		ID:        uuid.New().String(),
		Name:      "John Zerino",
		Email:     "john.doe@example.com",
		ImagePath: nil,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	})
}
