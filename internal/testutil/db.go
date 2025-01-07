package testutil

import (
	"database/sql"
	"os"
	"testing"
)

func NewTestDB(t *testing.T) *sql.DB {
	dbURL := os.Getenv("TEST_DB_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/app360_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatalf("could not connect to test database: %v", err)
	}

	// Limpar tabelas antes dos testes
	if _, err := db.Exec("TRUNCATE users CASCADE"); err != nil {
		t.Fatalf("could not truncate users table: %v", err)
	}

	return db
}
