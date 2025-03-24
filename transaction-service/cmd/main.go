package main

import (
	"log"
	"net/http"

	"bank-system/transaction-service/internal/application/ports/input"
	"bank-system/transaction-service/internal/infrastructure/adapters/input/rest"
	httpAdapter "bank-system/transaction-service/internal/infrastructure/adapters/output/http"
	"bank-system/transaction-service/internal/infrastructure/adapters/output/in_memory_db"
)

var transactionRepo = in_memory_db.NewInMemoryTransactionRepository()
var accountService = httpAdapter.NewAccountHttpAdapter("http://localhost:8081")

var processTransactionUseCase = input.NewCreateTransactionInputPort(transactionRepo, accountService)
var processTransactionController = rest.NewProcessTransactionController(processTransactionUseCase)

var listTransactionsUseCase = input.NewGetTransactionsInputPort(transactionRepo)
var listTransactionsController = rest.NewListTransactionsController(listTransactionsUseCase)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Account service is running!"))
}

func transactionProcess(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		processTransactionController.ProcessTransactionHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func transactionsList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		listTransactionsController.ListTransactionsHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/transactions", transactionProcess)
	mux.HandleFunc("/transactions/", transactionsList)

	log.Print("Starting Account Service on port 8082...")
	err := http.ListenAndServe(":8082", mux)
	log.Fatal(err)
}
