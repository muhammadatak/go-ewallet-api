package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type userRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.New("default", con),
	}
}

func (u *userRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	dataset := getExecutor(ctx, u.db).From("users").Where(goqu.C("id").Eq(id))
	var result domain.User
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	dataset := getExecutor(ctx, u.db).From("users").Where(goqu.C("username").Eq(username))
	var result domain.User
	_, err := dataset.ScanStructContext(ctx, &result)
	if err != nil {
		return domain.User{}, err
	}
	return result, nil
}

func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	executor := getExecutor(ctx, u.db).Insert("users").Rows(user).Returning("id").Executor()
	_, err := executor.ScanStructContext(ctx, user)
	return err
}

// Update implements [domain.UserRepository].
func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	executor := getExecutor(ctx, u.db).Update("users").Where(goqu.C("id").Eq(user.Id)).Set(user).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
