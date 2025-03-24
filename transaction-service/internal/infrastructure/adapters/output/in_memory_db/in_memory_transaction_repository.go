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

func (r *InMemoryTransactionRepository) FindByAccountID(accountID int64) ([]*domain.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*domain.Transaction

	for _, tx := range r.transactions {
		if tx.FromAccount == accountID || tx.ToAccount == accountID {
			result = append(result, tx)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("no transactions found for the account")
	}

	return result, nil
}

var _ output.TransactionOutputPort = (*InMemoryTransactionRepository)(nil)
