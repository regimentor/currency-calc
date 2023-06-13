package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/regimentor/currency-calc/internal"
	"github.com/regimentor/currency-calc/internal/api/http"
	currencyapi_com "github.com/regimentor/currency-calc/internal/currencyapi.com"
	"github.com/regimentor/currency-calc/internal/postgresql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	psqlUser := os.Getenv("POSTGRESQL_USERNAME")
	psqlPass := os.Getenv("POSTGRESQL_PASSWORD")
	psqlDb := os.Getenv("POSTGRESQL_DATABASE")
	currencyComApiKey := os.Getenv("CURRENCIES_COM_API_KEY")

	// TODO: pass a context
	poll, err := NewConnection(psqlUser, psqlPass, psqlDb)
	if err != nil {
		log.Fatalf("connection to database due err: %v", err)
	}

	userStorage := postgresql.NewUserStorage(poll)
	userRepository := internal.NewUserRepository(userStorage)

	currencyStorage := postgresql.NewCurrencyStorage(poll)
	currencyApiCom := currencyapi_com.NewCurrencyApiCom(currencyapi_com.ApiKey(currencyComApiKey))
	currencyRepository := internal.NewCurrencyRepository(currencyStorage, currencyApiCom)

	apiLogsStorage := postgresql.NewApiLogStorage(poll)
	apiLogsRepository := internal.NewApiLogsRepository(apiLogsStorage)

	httpServer := http.NewServer(userRepository, currencyRepository, apiLogsRepository)

	err = httpServer.Listen()
	if err != nil {
		log.Fatalf("create http server due err: %v", err)
	}

	log.Println(poll)
}
