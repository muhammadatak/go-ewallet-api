package sse

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/api"
)

type notificationSse struct {
	hub *dto.Hub
}

func NewNotification(app *fiber.App, authMidd fiber.Handler, hub *dto.Hub) {
	h := notificationSse{
		hub: hub,
	}

	app.Get("/sse/notification-stream", authMidd, h.StreamNotification)
}

func (n notificationSse) StreamNotification(ctx fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")

	userId, err := api.GetUserId(ctx)
	if err != nil {
		return err
	}

	n.hub.NotificationChannel[userId] = make(chan dto.NotificationData)

	ctx.RequestCtx().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("event %s\n"+
			"data: \n\n", "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush()
		for notification := range n.hub.NotificationChannel[userId] {
			data, _ := json.Marshal(notification)

			event = fmt.Sprintf("event %s\n"+
				"data: %s\n\n", "notification updated", data)

			_, _ = fmt.Fprint(w, event)
			_ = w.Flush()
		}
	})
	return nil
}
