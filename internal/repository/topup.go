package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type topUpRepository struct {
	db *goqu.Database
}

func NewTopUp(con *sql.DB) domain.TopUpRepository {
	return &topUpRepository{
		db: goqu.New("default", con),
	}
}

// FindById implements [domain.TopUpRepository].
func (t *topUpRepository) FindById(ctx context.Context, id string) (domain.TopUp, error) {
	dataset := getExecutor(ctx, t.db).From("topups").Where(goqu.C("id").Eq(id))
	var result domain.TopUp
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.TopUp{}, err
	}

	return result, nil
}

// Insert implements [domain.TopUpRepository].
func (t *topUpRepository) Insert(ctx context.Context, topup *domain.TopUp) error {
	executor := getExecutor(ctx, t.db).Insert("topups").Rows(topup).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements [domain.TopUpRepository].
func (t *topUpRepository) Update(ctx context.Context, topup *domain.TopUp) error {
	executor := getExecutor(ctx, t.db).Update("topups").Where(goqu.C("id").Eq(topup.Id)).
		Set(goqu.Record{
			"amount":   topup.Amount,
			"status":   topup.Status,
			"snap_url": topup.SnapUrl,
		}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
