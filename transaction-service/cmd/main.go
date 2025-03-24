package main

import (
	"log"
	"net/http"

	"bank-system/transaction-service/internal/application/ports/input"
	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/application/usecases"
	"bank-system/transaction-service/internal/infrastructure/adapters/input/rest"
	httpAdapter "bank-system/transaction-service/internal/infrastructure/adapters/output/http"
	"bank-system/transaction-service/internal/infrastructure/adapters/output/postgresql"
)

var (
	accountService               output.AccountOutputPort
	processTransactionUseCase    usecases.ProcessTransferUseCase
	processTransactionController *rest.ProcessTransactionController
	listTransactionsUseCase      usecases.GetTransactionsUseCase
	listTransactionsController   *rest.ListTransactionsController
)

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
	connStr := "postgres://postgres:David007@localhost:5432/banksystem?sslmode=disable"
	transactionRepo, err := postgresql.NewPostgresTransactionRepository(connStr)

	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}

	accountService = httpAdapter.NewAccountHttpAdapter("http://localhost:8081")

	processTransactionUseCase = input.NewCreateTransactionInputPort(transactionRepo, accountService)
	processTransactionController = rest.NewProcessTransactionController(processTransactionUseCase)

	listTransactionsUseCase = input.NewGetTransactionsInputPort(transactionRepo)
	listTransactionsController = rest.NewListTransactionsController(listTransactionsUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/transactions", transactionProcess)
	mux.HandleFunc("/transactions/", transactionsList)

	log.Print("Starting Account Service on port 8082...")

	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
