package input

import (
	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/application/usecases"
	"bank-system/transaction-service/internal/domain"
)

type GetTransactionsInputPort struct {
	repo output.TransactionOutputPort
}

func NewGetTransactionsInputPort(repo output.TransactionOutputPort) usecases.GetTransactionsUseCase {
	return &GetTransactionsInputPort{repo: repo}
}

func (g *GetTransactionsInputPort) Execute(accountID int64) ([]*domain.Transaction, error) {
	return g.repo.FindByAccountID(accountID)
}
