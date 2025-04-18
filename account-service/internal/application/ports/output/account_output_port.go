package output

import "bank-system/account-service/internal/domain"

type AccountOutputPort interface {
	Save(account *domain.Account) error
	FindByID(id int64) (*domain.Account, error)
}
