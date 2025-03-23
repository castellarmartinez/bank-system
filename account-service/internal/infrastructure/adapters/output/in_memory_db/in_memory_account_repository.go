package in_memory_db

import (
	"errors"
	"sync"

	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/domain"
)

type InMemoryAccountRepository struct {
	mu       sync.RWMutex
	accounts map[int64]*domain.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		accounts: make(map[int64]*domain.Account),
	}
}

func (r *InMemoryAccountRepository) Save(account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if account == nil {
		return errors.New("account cannot be null")
	}

	account.Id = int64(len(r.accounts) + 1)
	r.accounts[account.Id] = account

	return nil
}

func (r *InMemoryAccountRepository) FindByID(id int64) (*domain.Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, exists := r.accounts[id]

	if !exists {
		return nil, errors.New("account not found")
	}

	return account, nil
}

var _ output.AccountOutputPort = (*InMemoryAccountRepository)(nil)
