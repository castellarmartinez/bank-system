package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bank-system/account-service/internal/application/usecases"
)

type GetAccountBalanceController struct {
	useCase usecases.GetAccountBalanceUseCase
}

type BalanceResponse struct {
	Balance float64 `json:"balance"`
}

func NewGetAccountBalanceController(useCase usecases.GetAccountBalanceUseCase) *GetAccountBalanceController {
	return &GetAccountBalanceController{
		useCase: useCase,
	}
}

func (a *GetAccountBalanceController) GetAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/accounts/")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		http.Error(w, "invalid account id", http.StatusBadRequest)
		return
	}

	balance, err := a.useCase.Execute(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := BalanceResponse{
		Balance: balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
