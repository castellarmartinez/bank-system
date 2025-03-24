package domain

import (
	"errors"
	"time"
)

const (
	StatusSuccess = "success"
	StatusFailed  = "failed"
)

type Transaction struct {
	ID          int64
	FromAccount int64
	ToAccount   int64
	Amount      float64
	Timestamp   time.Time
	Status      string
}

func NewTransaction(id, from, to int64, amount float64) (*Transaction, error) {

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	if from == to {
		return nil, errors.New("cannot transfer money to the same account")
	}

	return &Transaction{
		ID:          id,
		FromAccount: from,
		ToAccount:   to,
		Amount:      amount,
		Timestamp:   time.Now().UTC(),
		Status:      StatusSuccess,
	}, nil
}
