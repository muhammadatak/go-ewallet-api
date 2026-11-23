package api

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/go-e-wallet/domain"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topUpService    domain.TopUpService
}

func NewMidtrans(app *fiber.App, midtransService domain.MidtransService, topUpService domain.TopUpService) {
	ma := midtransApi{
		midtransService: midtransService,
		topUpService:    topUpService,
	}

	app.Post("midtrans/payment-callback", ma.paymentHandlerNotification)
}

func (ma midtransApi) paymentHandlerNotification(ctx fiber.Ctx) error {
	var notificationPayload map[string]interface{}

	if err := ctx.Bind().Body(&notificationPayload); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return ctx.SendStatus(http.StatusBadRequest)
	}

	success, _ := ma.midtransService.VerifyPayment(ctx.Context(), orderId)
	if success {
		_ = ma.topUpService.ConfirmedTopUp(ctx.Context(), orderId)
		return ctx.SendStatus(http.StatusOK)
	}
	return ctx.SendStatus(http.StatusBadRequest)
}
