package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type loginLogRepository struct {
	db *goqu.Database
}

func NewLoginLog(con *sql.DB) domain.LoginLogRepository {
	return &loginLogRepository{
		db: goqu.New("default", con),
	}
}

// FindLastAuthorized implements [domain.LoginLogRepository].
func (l *loginLogRepository) FindLastAuthorized(ctx context.Context, userId int64) (domain.LoginLog, error) {
	dataset := getExecutor(ctx, l.db).From("login_log").Where(goqu.Ex{
		"user_id":       userId,
		"is_authorized": true,
	}).Order(goqu.I("id").Desc()).Limit(1)

	var result domain.LoginLog
	_, err := dataset.ScanStructContext(ctx, &result)

	if err != nil {
		return domain.LoginLog{}, err
	}
	return result, nil
}

// Insert implements [domain.LoginLogRepository].
func (l *loginLogRepository) Insert(ctx context.Context, login *domain.LoginLog) error {
	executor := l.db.Insert("login_log").Rows(goqu.Record{
		"user_id":       login.UserId,
		"is_authorized": login.IsAuthorized,
		"ip_address":    login.IpAddress,
		"timezone":      login.Timezone,
		"lat":           login.Lat,
		"lon":           login.Lon,
		"access_time":   login.AccessTime,
	}).Returning("id").Executor()

	_, err := executor.ScanStructContext(ctx, login)
	return err
}
