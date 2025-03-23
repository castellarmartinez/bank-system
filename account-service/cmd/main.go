package main

import (
	"log"
	"net/http"

	"bank-system/account-service/internal/application/ports/input"
	"bank-system/account-service/internal/infrastructure/adapters/input/rest"
	"bank-system/account-service/internal/infrastructure/adapters/output/in_memory_db"
)

var accountRepo = in_memory_db.NewInMemoryAccountRepository()

var createAccountUseCase = input.NewCreateAccountInputPort(accountRepo)
var createAccountController = rest.NewCreateAccountController(createAccountUseCase)

var getAccountBalanceUseCase = input.NewGetAccountBalanceInputPort(accountRepo)
var getAccountBalanceController = rest.NewGetAccountBalanceController(getAccountBalanceUseCase)

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
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/accounts", accountCreate)
	mux.HandleFunc("/accounts/", accountGetBalance)

	log.Print("Starting Account Service on port 8081...")
	err := http.ListenAndServe(":8081", mux)
	log.Fatal(err)
}
