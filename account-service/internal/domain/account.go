package domain

import (
	"errors"
)

type Account struct {
	Id      int64
	Name    string
	Balance float64
}

func NewAccount(id int64, name string, balance float64) (*Account, error) {
	if balance < 0 {
		return nil, errors.New("balance must be non-negative")
	}

	if name == "" {
		return nil, errors.New("account name is mandatory")
	}

	return &Account{
		Id:      id,
		Name:    name,
		Balance: float64(balance),
	}, nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}
