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

type userApi struct {
	userService domain.UserService
	fdsService  domain.FdsService
}

func NewUser(app *fiber.App, userService domain.UserService, fdsService domain.FdsService) {
	ua := userApi{
		userService: userService,
		fdsService:  fdsService,
	}

	app.Post("/login", ua.Login)
	app.Post("/register", ua.Register)
	app.Post("/validate-otp", ua.ValidateOTP)

}

func (ua userApi) Login(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.LoginRequest

	err := ctx.Bind().Body(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError(err.Error()))
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}
	res, err := ua.userService.Login(c, req)

	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}
	if !ua.fdsService.IsAuthorized(c, ctx.Get("X-FORWARDED-FOR"), res.UserId) {
		return ctx.SendStatus(http.StatusUnauthorized)
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ua userApi) Register(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.RegisterRequest
	err := ctx.Bind().Body(&req)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}

	res, err := ua.userService.Register(c, req)

	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ua userApi) ValidateOTP(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancel()

	var req dto.ValidateOTPReq
	err := ctx.Bind().Body(&req)

	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}

	err = ua.userService.ValidateOTP(c, req)

	if err != nil {
		return ctx.Status(errorStatus(err)).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess("ok"))

}
