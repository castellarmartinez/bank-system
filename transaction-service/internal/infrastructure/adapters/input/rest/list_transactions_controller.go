package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"bank-system/transaction-service/internal/application/usecases"
)

type ListTransactionsController struct {
	useCase usecases.GetTransactionsUseCase
}

type TransactionsResponse struct {
	Status      string  `json:"status"`
	ID          int64   `json:"transaction_id"`
	FromAccount int64   `json:"from_account"`
	ToAccount   int64   `json:"to_account"`
	Amount      float64 `json:"monto"`
	Timestamp   string  `json:"fecha"`
}

func NewListTransactionsController(listTransactionsUseCase usecases.GetTransactionsUseCase) *ListTransactionsController {
	return &ListTransactionsController{
		useCase: listTransactionsUseCase,
	}
}

func (l *ListTransactionsController) ListTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/transactions/")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	txs, err := l.useCase.Execute(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response []TransactionsResponse

	for _, tx := range txs {
		response = append(response, TransactionsResponse{
			tx.Status,
			tx.ID,
			tx.FromAccount,
			tx.ToAccount,
			tx.Amount,
			tx.Timestamp.Format(time.RFC3339),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
