package in_memory_db

import (
	"errors"
	"sync"

	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/domain"
)

type InMemoryAccountRepository struct {
	mu    sync.RWMutex
	store map[int64]*domain.Account
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		store: make(map[int64]*domain.Account),
	}
}

func (r *InMemoryAccountRepository) Save(account *domain.Account) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if account == nil {
		return errors.New("account cannot be null")
	}

	r.store[account.Id] = account
	return nil
}

var _ output.AccountOutputPort = (*InMemoryAccountRepository)(nil)
