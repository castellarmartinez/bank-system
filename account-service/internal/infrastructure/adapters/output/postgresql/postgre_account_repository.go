package postgres

import (
	"bank-system/account-service/internal/application/ports/output"
	"bank-system/account-service/internal/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresAccountRepository struct {
	db *sql.DB
}

func NewPostgresAccountRepository(connStr string) (*PostgresAccountRepository, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres ping failed: %v", err)
	}

	return &PostgresAccountRepository{db: db}, nil
}

func (r *PostgresAccountRepository) Save(account *domain.Account) error {
	ctx := context.Background()

	query := `
        INSERT INTO accounts
        (name, balance)
        VALUES ($1, $2)
				RETURNING id
    `

	err := r.db.QueryRowContext(ctx, query,
		account.Name,
		account.Balance,
	).Scan(&account.ID)

	return err
}

func (r *PostgresAccountRepository) FindByID(id int64) (*domain.Account, error) {
	ctx := context.Background()
	query := `
        SELECT id, name, balance
        FROM accounts
        WHERE id = $1
    `

	row := r.db.QueryRowContext(ctx, query, id)

	var account domain.Account

	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.Balance,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}

		return nil, fmt.Errorf("failed to scan account: %v", err)
	}

	return &account, nil
}

var _ output.AccountOutputPort = (*PostgresAccountRepository)(nil)
