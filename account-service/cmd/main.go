package main

import (
	"log"
	"net/http"

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
	connStr := "postgres://postgres:David007@localhost:5432/banksystem?sslmode=disable"
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

	log.Print("Starting Account Service on port 8081...")

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatal(err)
	}
}
