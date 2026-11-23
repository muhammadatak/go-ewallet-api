package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/util"
)

type topUpApi struct {
	topUpService domain.TopUpService
}

func NewTopUpApi(app *fiber.App, topUpService domain.TopUpService, authMidd fiber.Handler) {
	ta := topUpApi{
		topUpService: topUpService,
	}

	app.Post("/topup/initialize", authMidd, ta.InitializeTopUp)
}

func (t topUpApi) InitializeTopUp(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TopUpReq

	err := ctx.Bind().Body(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}

	userId, err := GetUserId(ctx)
	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}

	req.UserId = userId

	res, err := t.topUpService.InitializedTopUp(c, req)
	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))

}
