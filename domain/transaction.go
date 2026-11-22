package domain

import (
	"context"
	"time"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type Transaction struct {
	Id                  int64     `db:"id" goqu:"skipinsert,skipupdate"`
	AccountId           int64     `db:"account_id"`
	SofNumber           string    `db:"sof_number"`
	DofNumber           string    `db:"dof_number"`
	TransactionType     string    `db:"transaction_type"`
	Amount              float64   `db:"amount"`
	TransactionDateTime time.Time `db:"transaction_datetime"`
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction *Transaction) error
}

type TransactionService interface {
	TransferInquiry(ctx context.Context, req dto.TransferInquiryReq, userId int64) (dto.TransferInquiryRes, error)
	TransferExecute(ctx context.Context, req dto.TransferExecuteReq, userId int64) error
}
