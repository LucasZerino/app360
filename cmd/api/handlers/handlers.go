package handlers

import (
	"app360/internal/auth"
	"app360/internal/store"
)

type Handlers struct {
	User *UserHandler
	Auth *AuthHandler
}

func NewHandlers(store *store.Storage, jwt *auth.JWTManager) *Handlers {
	return &Handlers{
		User: NewUserHandler(store),
		Auth: NewAuthHandler(store, jwt),
	}
}
