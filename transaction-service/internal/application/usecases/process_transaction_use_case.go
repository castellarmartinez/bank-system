package usecases

import "bank-system/transaction-service/internal/domain"

type ProcessTransferUseCase interface {
	Execute(fromAccount int64, toAccount int64, amount float64) (*domain.Transaction, error)
}
