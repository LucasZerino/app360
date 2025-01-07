package main

import (
	"log"

	"app360/cmd/api/handlers"
	"app360/internal/auth"
	"app360/internal/db"
	"app360/internal/env"
	"app360/internal/store"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	config   config
	store    *store.Storage
	handlers *handlers.Handlers
	jwt      *auth.JWTManager
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Arquivo .env não encontrado: %v", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:postgres@localhost:5432/app360?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	store := store.NewStorage(db)
	jwtManager := auth.NewJWTManager(env.GetString("JWT_SECRET", "seu-segredo-aqui"))

	app := &application{
		config:   cfg,
		store:    store,
		handlers: handlers.NewHandlers(store, jwtManager),
		jwt:      jwtManager,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
