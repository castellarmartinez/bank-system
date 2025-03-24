package postgresql

import (
	"bank-system/transaction-service/internal/application/ports/output"
	"bank-system/transaction-service/internal/domain"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresTransactionRepository struct {
	db *sql.DB
}

func NewPostgresTransactionRepository(connStr string) (*PostgresTransactionRepository, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres ping failed: %v", err)
	}

	return &PostgresTransactionRepository{db: db}, nil
}

func (r *PostgresTransactionRepository) Save(tx *domain.Transaction) error {
	ctx := context.Background()
	query := `
        INSERT INTO transactions
            (from_account, to_account, amount, status, timestamp)
        VALUES ($1, $2, $3, $4, $5)
				RETURNING id
    `

	err := r.db.QueryRowContext(ctx, query,
		tx.FromAccount,
		tx.ToAccount,
		tx.Amount,
		tx.Status,
		tx.Timestamp,
	).Scan(&tx.ID)

	return err
}

func (r *PostgresTransactionRepository) FindByAccountID(id int64) ([]*domain.Transaction, error) {
	ctx := context.Background()
	query := `
        SELECT id, from_account, to_account, amount, status, timestamp
        FROM transactions
        WHERE from_account = $1 OR to_account = $1
    `

	rows, err := r.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var transactions []*domain.Transaction

	for rows.Next() {
		var t domain.Transaction

		err := rows.Scan(
			&t.ID,
			&t.FromAccount,
			&t.ToAccount,
			&t.Amount,
			&t.Status,
			&t.Timestamp,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, &t)
	}

	return transactions, nil
}

var _ output.TransactionOutputPort = (*PostgresTransactionRepository)(nil)
