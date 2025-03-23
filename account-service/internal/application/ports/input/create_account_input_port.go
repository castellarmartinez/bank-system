package input

import (
	"sync/atomic"

	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/application/usecases"
	"bank-system/account-service/internal/domain"
)

type AccountInputPort struct {
	repo      output.AccountOutputPort
	idCounter int64
}

func NewCreateAccountInputPort(repo output.AccountOutputPort) usecases.CreateAccountUseCase {
	return &AccountInputPort{repo: repo, idCounter: 0}
}

func (a *AccountInputPort) nextID() int64 {
	return atomic.AddInt64(&a.idCounter, 1)
}

func (a *AccountInputPort) Execute(name string, balance float64) (*domain.Account, error) {
	account, err := domain.NewAccount(a.nextID(), name, balance)

	if err != nil {
		return nil, err
	}

	if err := a.repo.Save(account); err != nil {
		return nil, err
	}

	return account, nil
}
