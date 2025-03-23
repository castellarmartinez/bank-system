package usecases

import "bank-system/account-service/internal/domain"

type CreateAccountUseCase interface {
	CreateAccount(name string, balance float64) (*domain.Account, error)
}
