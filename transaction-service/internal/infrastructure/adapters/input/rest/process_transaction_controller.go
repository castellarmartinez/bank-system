package rest

import (
	"encoding/json"
	"net/http"

	"bank-system/transaction-service/internal/application/usecases"
)

type ProcessTransactionController struct {
	useCase usecases.ProcessTransferUseCase
}

type TransactionResponse struct {
	Status string `json:"status"`
	ID     int64  `json:"transaction_id"`
}

func NewProcessTransactionController(processTransactionUseCase usecases.ProcessTransferUseCase) *ProcessTransactionController {
	return &ProcessTransactionController{
		useCase: processTransactionUseCase,
	}
}

func (a *ProcessTransactionController) ProcessTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FromAccount int64   `json:"from_account"`
		ToAccount   int64   `json:"to_account"`
		Amount      float64 `json:"monto"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	tx, err := a.useCase.Execute(req.FromAccount, req.ToAccount, req.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := TransactionResponse{
		tx.Status,
		tx.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
