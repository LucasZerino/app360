package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	User interface {
		Create(ctx context.Context, user *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		User: &UserStore{db},
	}
}
