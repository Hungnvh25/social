package main

import (
	"log"

	"github.com/Hungnvh25/social/internal/db"
	"github.com/Hungnvh25/social/internal/env"
	"github.com/Hungnvh25/social/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Configure application settings
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// Initialize database
	database, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.CloseDB(database)
	log.Println("Connected to the database")

	// Create a new store instance
	store := store.NewStorage(database)

	// Initialize the application
	app := &application{
		config: cfg,
		store:  store,
	}

	// Mount routes and run the application
	mux := app.mount()
	log.Fatal(app.run(mux))
}
