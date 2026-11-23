package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
)

type notificationApi struct {
	notificationService domain.NotificationService
}

func NewNotification(app *fiber.App, notificationService domain.NotificationService, jwtMidd fiber.Handler) {
	na := notificationApi{
		notificationService: notificationService,
	}

	app.Get("notification", jwtMidd, na.Index)
}

func (na notificationApi) Index(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	userId, err := GetUserId(ctx)
	if err != nil {
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	res, err := na.notificationService.FindByUserId(c, userId)

	if err != nil {
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))

}
