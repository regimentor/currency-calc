package main

import (
	"github.com/joho/godotenv"
	"github.com/regimentor/currency-calc/internal/api/http"
	"github.com/regimentor/currency-calc/internal/postgresql"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	psqlUser := os.Getenv("POSTGRESQL_USERNAME")
	psqlPass := os.Getenv("POSTGRESQL_PASSWORD")
	psqlDb := os.Getenv("POSTGRESQL_DATABASE")

	// TODO: pass a context
	poll, err := postgresql.NewConnection(psqlUser, psqlPass, psqlDb)
	if err != nil {
		log.Fatalf("connection to database due err: %v", err)
	}

	httpServer := http.NewServer()
	err = httpServer.Listen()
	if err != nil {
		log.Fatalf("create http server due err: %v", err)
	}

	log.Println(poll)
}