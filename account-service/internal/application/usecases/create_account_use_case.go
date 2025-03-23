package usecases

import "bank-system/account-service/internal/domain"

type CreateAccountUseCase interface {
	Execute(name string, balance float64) (*domain.Account, error)
}
