package output

import "bank-system/transaction-service/internal/domain"

type TransactionOutputPort interface {
	Save(transaction *domain.Transaction) error
	FindByAccountID(accountID int64) ([]*domain.Transaction, error)
}
