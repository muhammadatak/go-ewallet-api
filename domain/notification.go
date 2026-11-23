package domain

import (
	"context"
	"time"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type Notification struct {
	Id        int64     `db:"id" goqu:"skipinsert,skipupdate"`
	UserId    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	Status    int8      `db:"status"`
	IsRead    int8      `db:"is_read"`
	CreatedAt time.Time `db:"created_at"`
}

type NotificationRepository interface {
	FindByUserId(ctx context.Context, id int64) ([]Notification, error)
	Insert(ctx context.Context, notification *Notification) error
	Update(ctx context.Context, notification *Notification) error
}

type NotificationService interface {
	FindByUserId(ctx context.Context, id int64) ([]dto.NotificationData, error)
	Insert(ctx context.Context, userId int64, code string, data map[string]string) error
}
