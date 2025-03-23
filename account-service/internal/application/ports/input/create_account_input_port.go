package input

import (
	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/application/usecases"
	"bank-system/account-service/internal/domain"
)

type CreateAccountInputPort struct {
	repo output.AccountOutputPort
}

func NewCreateAccountInputPort(repo output.AccountOutputPort) usecases.CreateAccountUseCase {
	return &CreateAccountInputPort{repo: repo}
}

func (c *CreateAccountInputPort) Execute(name string, balance float64) (*domain.Account, error) {
	account, err := domain.NewAccount(0, name, balance) // id is then set in the repository

	if err != nil {
		return nil, err
	}

	if err := c.repo.Save(account); err != nil {
		return nil, err
	}

	return account, nil
}
