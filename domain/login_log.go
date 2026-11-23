package domain

import (
	"context"
	"time"
)

type LoginLog struct {
	Id           int64     `db:"id"`
	UserId       int64     `db:"user_id"`
	IsAuthorized bool      `db:"is_authorized"`
	IpAddress    string    `db:"ip_address"`
	Timezone     string    `db:"timezone"`
	Lat          float64   `db:"lat"`
	Lon          float64   `db:"lon"`
	AccessTime   time.Time `db:"access_time"`
}

type LoginLogRepository interface {
	FindLastAuthorized(ctx context.Context, userId int64) (LoginLog, error)
	Insert(ctx context.Context, login *LoginLog) error
}
