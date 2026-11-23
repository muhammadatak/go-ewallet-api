package main

import (
	"log"
	"net/http"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/api"
	"github.com/muhammadatak/go-e-wallet/internal/config"
	"github.com/muhammadatak/go-e-wallet/internal/connection"
	"github.com/muhammadatak/go-e-wallet/internal/repository"
	"github.com/muhammadatak/go-e-wallet/internal/service"
	"github.com/muhammadatak/go-e-wallet/internal/sse"
)

func main() {
	cnf := config.Get()

	jwtMidd := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("endpoint perlu token, silakan login dulu"))
		},
	})
	dbConn := connection.GetDatabase(cnf.Database)
	cacheConn := repository.NewRedisClient(cnf)
	app := fiber.New()

	hub := &dto.Hub{
		NotificationChannel: map[int64]chan dto.NotificationData{},
	}

	//repos
	userRepos := repository.NewUser(dbConn)
	accountRepos := repository.NewAccount(dbConn)
	transactionRepos := repository.NewTransaction(dbConn)
	notificationRepos := repository.NewNotification(dbConn)
	templateRepos := repository.NewTemplate(dbConn)

	//service
	emailService := service.NewEmail(cnf)
	userService := service.NewUser(userRepos, cacheConn, cnf, emailService)
	notificationService := service.NewNotification(notificationRepos, userRepos, templateRepos, hub)

	transactionService := service.NewTransaction(accountRepos, transactionRepos, cacheConn, notificationService)

	//api
	api.NewUser(app, userService)
	api.NewTransfer(app, transactionService, jwtMidd)
	api.NewNotification(app, notificationService, jwtMidd)

	sse.NewNotification(app, jwtMidd, hub)
	log.Fatal(app.Listen(cnf.Server.Host + ":" + cnf.Server.Port))

}
