package input

import (
	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/application/usecases"
)

type GetAccountBalanceInputPort struct {
	repo output.AccountOutputPort
}

func NewGetAccountBalanceInputPort(repo output.AccountOutputPort) usecases.GetAccountBalanceUseCase {
	return &GetAccountBalanceInputPort{repo: repo}
}

func (a *GetAccountBalanceInputPort) Execute(id int64) (float64, error) {
	account, err := a.repo.FindByID(id)

	if err != nil {
		return 0, err
	}

	return account.GetBalance(), nil
}
