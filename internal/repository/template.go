package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type templateRepository struct {
	db *goqu.Database
}

func NewTemplate(con *sql.DB) domain.TemplateRepository {
	return &templateRepository{
		db: goqu.New("default", con),
	}
}

// FindByCode implements [domain.TemplateRepository].
func (t *templateRepository) FindByCode(ctx context.Context, code string) (domain.Template, error) {
	dataset := t.db.From("templates").Where(goqu.C("code").Eq(code))
	var result domain.Template
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.Template{}, err
	}

	return result, nil
}
