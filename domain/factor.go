package domain

import (
	"context"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type Factor struct {
	Id     int64  `db:"id"`
	UserId int64  `db:"user_id"`
	PIN    string `db:"pin"`
}

type FactorRepository interface {
	FindByUser(ctx context.Context, userId int64) (Factor, error)
}

type FactorService interface {
	ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error
}
