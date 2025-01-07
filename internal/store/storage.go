package store

import "database/sql"

// Storage contém todos os stores da aplicação
type Storage struct {
	User *UserStore
}

// NewStorage inicializa todos os stores
func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		User: NewUserStore(db),
	}
}
