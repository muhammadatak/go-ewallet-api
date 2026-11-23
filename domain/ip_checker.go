package domain

import (
	"context"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type IpCheckerService interface {
	Query(ctx context.Context, ip string) (dto.IpChecker, error)
}
