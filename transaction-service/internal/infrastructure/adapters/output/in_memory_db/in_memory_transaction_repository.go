package in_memory_db

import (
	"errors"
	"sync"

	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/domain"
)

type InMemoryTransactionRepository struct {
	mu           sync.RWMutex
	transactions map[int64]*domain.Transaction
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[int64]*domain.Transaction),
	}
}

func (r *InMemoryTransactionRepository) Save(transaction *domain.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if transaction == nil {
		return errors.New("transaction cannot be null")
	}

	transaction.ID = int64(len(r.transactions) + 1)
	r.transactions[transaction.ID] = transaction

	return nil
}

var _ output.TransactionOutputPort = (*InMemoryTransactionRepository)(nil)
