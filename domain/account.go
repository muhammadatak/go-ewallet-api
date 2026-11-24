package domain

import "context"

type Account struct {
	Id            int64   `db:"id" goqu:"skipinsert,skipupdate"`
	UserId        int64   `db:"user_id"`
	AccountNumber string  `db:"account_number"`
	Balance       float64 `db:"balance"`
}

type AccountRepository interface {
	FindByUserId(ctx context.Context, id int64) (Account, error)
	FindByAccountNumber(ctx context.Context, accNumber string) (Account, error)
	Update(ctx context.Context, account *Account) error
	Insert(ctx context.Context, account *Account) error
}
