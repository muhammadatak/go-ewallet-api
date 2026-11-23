package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type notificationRepository struct {
	db *goqu.Database
}

func NewNotification(con *sql.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: goqu.New("default", con),
	}
}

func (n *notificationRepository) FindByUserId(ctx context.Context, id int64) ([]domain.Notification, error) {
	dataset := getExecutor(ctx, n.db).From("notifications").
		Where(goqu.C("user_id").Eq(id)).Order(goqu.I("created_at").Desc()).Limit(15)
	var result []domain.Notification

	err := dataset.ScanStructsContext(ctx, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Insert implements [domain.NotificationRepository].
func (n *notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	executor := getExecutor(ctx, n.db).Insert("notifications").Rows(notification).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements [domain.NotificationRepository].
func (n *notificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	executor := getExecutor(ctx, n.db).Update("notifications").Where(goqu.C("id").Eq(notification.Id)).
		Set(notification).Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
