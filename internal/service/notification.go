package service

import (
	"bytes"
	"context"
	"html/template"
	"time"

	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
)

type notificationService struct {
	notificationRepository domain.NotificationRepository
	userRepository         domain.UserRepository
	templateRepository     domain.TemplateRepository
	hub                    *dto.Hub
}

func NewNotification(notificationRepository domain.NotificationRepository, userRepository domain.UserRepository, templateRepository domain.TemplateRepository, hub *dto.Hub) domain.NotificationService {
	return &notificationService{
		notificationRepository: notificationRepository,
		userRepository:         userRepository,
		templateRepository:     templateRepository,
		hub:                    hub,
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

// Insert implements [domain.NotificationService].
func (n *notificationService) Insert(ctx context.Context, userId int64, code string, data map[string]string) error {
	tmpl, err := n.templateRepository.FindByCode(ctx, code)

	if err != nil {
		return err
	}

	if tmpl == (domain.Template{}) {
		return domain.ErrNotFound
	}

	body := new(bytes.Buffer)
	t := template.Must(template.New("notif").Parse(tmpl.Body))
	err = t.Execute(body, data)
	if err != nil {
		return err
	}

	notification := domain.Notification{
		UserId:    userId,
		Title:     tmpl.Title,
		Body:      body.String(),
		Status:    1,
		IsRead:    0,
		CreatedAt: time.Now(),
	}

	err = n.notificationRepository.Insert(ctx, &notification)
	if err != nil {
		return err
	}

	if channel, ok := n.hub.NotificationChannel[userId]; ok {
		channel <- dto.NotificationData{
			Id:        notification.Id,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.Status,
			CreatedAt: notification.CreatedAt,
		}
	}
	return nil
}
