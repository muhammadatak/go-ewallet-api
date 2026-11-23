package service

import (
	"context"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
	"golang.org/x/crypto/bcrypt"
)

type factorService struct {
	factorRepository domain.FactorRepository
}

func NewFactor(factorRepository domain.FactorRepository) domain.FactorService {
	return &factorService{
		factorRepository: factorRepository,
	}

}

// ValidatePIN implements [domain.FactorService].
func (f *factorService) ValidatePIN(ctx context.Context, req dto.ValidatePinReq) error {
	factor, err := f.factorRepository.FindByUser(ctx, req.UserId)

	if err != nil {
		return err
	}

	if factor == (domain.Factor{}) {
		return domain.ErrNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(factor.PIN), []byte(req.PIN))
	if err != nil {
		return err
	}

	return nil
}
