package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"bank-system/account-service/internal/application/ports/input"
	"bank-system/account-service/internal/application/usecases"
	"bank-system/account-service/internal/infrastructure/adapters/input/rest"
	postgres "bank-system/account-service/internal/infrastructure/adapters/output/postgresql"
)

var (
	createAccountUseCase    usecases.CreateAccountUseCase
	createAccountController *rest.CreateAccountController

	getAccountBalanceUseCase    usecases.GetAccountBalanceUseCase
	getAccountBalanceController *rest.GetAccountBalanceController
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Account service is running!"))
}

func accountCreate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createAccountController.CreateAccountHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func accountGetBalance(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAccountBalanceController.GetAccountBalanceHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	dbHost := os.Getenv("DB_HOST")
	accountServicePort := os.Getenv("ACCOUNT_SERVICE_PORT")

	if dbHost == "" {
		dbHost = "localhost"
	}

	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
	accountRepo, err := postgres.NewPostgresAccountRepository(connStr)

	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}

	createAccountUseCase = input.NewCreateAccountInputPort(accountRepo)
	createAccountController = rest.NewCreateAccountController(createAccountUseCase)

	getAccountBalanceUseCase = input.NewGetAccountBalanceInputPort(accountRepo)
	getAccountBalanceController = rest.NewGetAccountBalanceController(getAccountBalanceUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/accounts", accountCreate)
	mux.HandleFunc("/accounts/", accountGetBalance)

	log.Print("Starting Account Service on port " + accountServicePort)

	if err := http.ListenAndServe(":"+accountServicePort, mux); err != nil {
		log.Fatal(err)
	}
}
