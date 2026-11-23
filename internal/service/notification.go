package service

import (
	"context"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	userRepository         domain.UserRepository
}

func NewNotification(notificationRepository domain.NotificationRepository, userRepository domain.UserRepository) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		userRepository:         userRepository,
	}
}

// FindByUserId implements [domain.NotificationService].
func (n *notificationService) FindByUserId(ctx context.Context, id int64) ([]dto.NotificationData, error) {
	user, err := n.userRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == (domain.User{}) {
		return nil, domain.ErrNotFound
	}

	notifications, err := n.notificationRepository.FindByUserId(ctx, user.Id)

	if err != nil {
		return nil, err
	}

	notificationData := make([]dto.NotificationData, 0)
	for _, v := range notifications {
		notificationData = append(notificationData, dto.NotificationData{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			Status:    v.Status,
			IsRead:    v.IsRead,
			CreatedAt: v.CreatedAt,
		})
	}

	return notificationData, nil

}
