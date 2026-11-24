package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type accountRepository struct {
	db *goqu.Database
}

// Insert implements [domain.AccountRepository].

func NewAccount(con *sql.DB) domain.AccountRepository {
	return &accountRepository{
		db: goqu.New("default", con),
	}
}

func (a *accountRepository) FindByAccountNumber(ctx context.Context, accNumber string) (domain.Account, error) {
	dataset := getExecutor(ctx, a.db).From("accounts").Where(goqu.C("account_number").Eq(accNumber))
	var result domain.Account
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.Account{}, err
	}

	return result, nil
}

func (a *accountRepository) FindByUserId(ctx context.Context, userId int64) (domain.Account, error) {
	dataset := getExecutor(ctx, a.db).From("accounts").Where(goqu.C("user_id").Eq(userId))
	var result domain.Account
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.Account{}, err
	}

	return result, nil
}

func (a *accountRepository) Update(ctx context.Context, account *domain.Account) error {
	dataset := getExecutor(ctx, a.db).Update("accounts").
		Where(goqu.C("id").Eq(account.Id)).Set(goqu.Record{"balance": account.Balance}).Executor()
	_, err := dataset.ExecContext(ctx)

	return err
}

func (a *accountRepository) Insert(ctx context.Context, account *domain.Account) error {
	executor := getExecutor(ctx, a.db).Insert("accounts").Rows(account).Returning("id").Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
