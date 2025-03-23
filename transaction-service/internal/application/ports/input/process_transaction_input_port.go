package input

import (
	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/application/usecases"
	"bank-system/transaction-service/internal/domain"
	"errors"
)

type CreateTransactionInputPort struct {
	repo           output.TransactionOutputPort
	accountService output.AccountOutputPort
}

func NewCreateTransactionInputPort(repo output.TransactionOutputPort, accountService output.AccountOutputPort) usecases.ProcessTransferUseCase {
	return &CreateTransactionInputPort{
		repo:           repo,
		accountService: accountService,
	}
}

func (c *CreateTransactionInputPort) Execute(fromAccount, toAccount int64, amount float64) (*domain.Transaction, error) {
	fromBalance, err := c.accountService.GetBalance(fromAccount)

	if err != nil {
		return nil, errors.New("from account does not exist or is unreachable")
	}

	_, err = c.accountService.GetBalance(toAccount)

	if err != nil {
		return nil, errors.New("to account does not exist or is unreachable")
	}

	if fromBalance < amount {
		return nil, errors.New("insufficient balance")
	}

	tx, err := domain.NewTransaction(0, fromAccount, toAccount, amount)

	if err != nil {
		return nil, err
	}

	if err := c.repo.Save(tx); err != nil {
		return nil, err
	}

	return tx, nil
}
