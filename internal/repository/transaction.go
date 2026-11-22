package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type transactionRepository struct {
	db *goqu.Database
}

func NewTransaction(con *sql.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: goqu.New("default", con),
	}
}

func (t *transactionRepository) Insert(ctx context.Context, transaction *domain.Transaction) error {
	executor := getExecutor(ctx, t.db).Insert("transactions").Rows(transaction).Returning("id").Executor()
	_, err := executor.ExecContext(ctx)

	return err
}
