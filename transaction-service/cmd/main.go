package main

import (
	"log"
	"net/http"
	"os"

	"bank-system/transaction-service/internal/application/ports/input"
	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/application/usecases"
	"bank-system/transaction-service/internal/infrastructure/adapters/input/rest"
	httpAdapter "bank-system/transaction-service/internal/infrastructure/adapters/output/http"
	"bank-system/transaction-service/internal/infrastructure/adapters/output/postgresql"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	host := os.Getenv("HOST")
	serverPort := os.Getenv("SERVER_PORT")
	accountServicePort := os.Getenv("ACCOUT_SERVICE_PORT")

	connStr := "postgres://" + dbUser + ":" + dbPassword + "@" + host + ":" + dbPort + "/" + dbName + "?sslmode=" + dbSSLMode
	transactionRepo, err := postgresql.NewPostgresTransactionRepository(connStr)

	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}

	accountService = httpAdapter.NewAccountHttpAdapter("http://" + host + ":" + accountServicePort)

	processTransactionUseCase = input.NewCreateTransactionInputPort(transactionRepo, accountService)
	processTransactionController = rest.NewProcessTransactionController(processTransactionUseCase)

	listTransactionsUseCase = input.NewGetTransactionsInputPort(transactionRepo)
	listTransactionsController = rest.NewListTransactionsController(listTransactionsUseCase)

	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/transactions", transactionProcess)
	mux.HandleFunc("/transactions/", transactionsList)

	log.Print("Starting Account Service on port " + serverPort)

	if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
		log.Fatal(err)
	}
}
