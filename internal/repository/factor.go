package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type factorRepository struct {
	db *goqu.Database
}

func NewFactor(con *sql.DB) domain.FactorRepository {
	return &factorRepository{
		db: goqu.New("default", con),
	}
}

// FindByUser implements [domain.FactorRepository].
func (f *factorRepository) FindByUser(ctx context.Context, userId int64) (domain.Factor, error) {
	dataset := getExecutor(ctx, f.db).From("factors").Where(goqu.C("user_id").Eq(userId))
	var result domain.Factor
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.Factor{}, err
	}

	return result, nil
}
