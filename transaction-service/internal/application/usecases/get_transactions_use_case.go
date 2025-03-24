package usecases

import "bank-system/transaction-service/internal/domain"

type GetTransactionsUseCase interface {
	Execute(accountID int64) ([]*domain.Transaction, error)
}
