package api

import (
	"errors"
	"net/http"
	"strconv"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"

	"github.com/muhammadatak/go-e-wallet/domain"
)

func errorStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, domain.ErrAuth):
		return http.StatusUnauthorized
	case errors.Is(err, domain.ErrInsufficientBalance),
		errors.Is(err, domain.ErrInvalidOTP),
		errors.Is(err, domain.ErrEmailNotVerified):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidToken):
		return http.StatusUnauthorized
	case errors.Is(err, domain.ErrUsernameExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func GetUserId(ctx fiber.Ctx) (int64, error) {
	token := jwtware.FromContext(ctx)
	if token == nil {
		return 0, domain.ErrInvalidToken
	}

	sub, err := token.Claims.GetSubject()
	if err != nil || sub == "" {
		return 0, domain.ErrInvalidToken
	}

	userId, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, domain.ErrInvalidToken
	}
	return userId, nil
}
