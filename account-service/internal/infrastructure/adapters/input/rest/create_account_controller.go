package rest

import (
	"encoding/json"
	"net/http"

	"bank-system/account-service/internal/application/usecases"
)

type CreateAccountController struct {
	useCase usecases.CreateAccountUseCase
}

func NewCreateAccountController(createAccountUseCase usecases.CreateAccountUseCase) *CreateAccountController {
	return &CreateAccountController{
		useCase: createAccountUseCase,
	}
}

func (a *CreateAccountController) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name    string  `json:"name"`
		Balance float64 `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	account, err := a.useCase.Execute(req.Name, req.Balance)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
