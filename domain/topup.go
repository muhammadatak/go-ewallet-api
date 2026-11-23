package domain

import (
	"context"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type TopUp struct {
	Id      string  `db:"id"`
	UserId  int64   `db:"user_id"`
	Status  int8    `db:"status"`
	Amount  float64 `db:"amount"`
	SnapUrl string  `db:"snap_url"`
}

type TopUpRepository interface {
	FindById(ctx context.Context, id string) (TopUp, error)
	Insert(ctx context.Context, t *TopUp) error
	Update(ctx context.Context, t *TopUp) error
}

type TopUpService interface {
	ConfirmedTopUp(ctx context.Context, id string) error
	InitializedTopUp(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error)
}
