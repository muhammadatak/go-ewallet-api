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

type transferApi struct {
	transactionService domain.TransactionService
}

func NewTransfer(app *fiber.App, transactionService domain.TransactionService, jwtMidd fiber.Handler) {
	ta := transferApi{
		transactionService: transactionService,
	}

	transfer := app.Group("/transfer", jwtMidd)

	transfer.Post("/inquiry", ta.TransferInquiry)
	transfer.Post("/execute", ta.TransferExecute)

}

func (ta transferApi) TransferInquiry(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TransferInquiryReq

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
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	res, err := ta.transactionService.TransferInquiry(c, req, userId)

	if err != nil {
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (ta transferApi) TransferExecute(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.TransferExecuteReq

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
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	err = ta.transactionService.TransferExecute(c, req, userId)

	if err != nil {
		return ctx.Status(errorStatus(err)).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusCreated)
}
