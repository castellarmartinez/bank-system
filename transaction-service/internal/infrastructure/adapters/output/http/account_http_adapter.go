package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"bank-system/transaction-service/internal/application/ports/output"
)

type AccountHttpAdapter struct {
	accountServiceURL string
}

func NewAccountHttpAdapter(url string) output.AccountOutputPort {
	return &AccountHttpAdapter{accountServiceURL: url}
}

func (a *AccountHttpAdapter) GetBalance(accountID int64) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("%s/accounts/%d", a.accountServiceURL, accountID))

	if err != nil {
		return 0, errors.New("failed to contact account service")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, errors.New("account not found")
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected response from account service: %d", resp.StatusCode)
	}

	var response struct {
		Balance float64 `json:"saldo"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, errors.New("failed to parse response")
	}

	return response.Balance, nil
}
