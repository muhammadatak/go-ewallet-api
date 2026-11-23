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
	topUpRepos := repository.NewTopUp(dbConn)
	factorRepos := repository.NewFactor(dbConn)
	loginLogRepos := repository.NewLoginLog(dbConn)

	//service
	ipCheckerService := service.NewIpChecker()
	fdsService := service.NewFds(ipCheckerService, loginLogRepos)
	emailService := service.NewEmail(cnf)
	factorService := service.NewFactor(factorRepos)
	userService := service.NewUser(userRepos, cacheConn, cnf, emailService)
	notificationService := service.NewNotification(notificationRepos, userRepos, templateRepos, hub)
	transactionService := service.NewTransaction(accountRepos, transactionRepos, cacheConn, notificationService)
	midtransService := service.NewMidtrans(cnf)
	topUpService := service.NewTopUp(notificationService, midtransService, topUpRepos, accountRepos, transactionRepos)
	//api
	api.NewUser(app, userService, fdsService)
	api.NewTransfer(app, transactionService, factorService, jwtMidd)
	api.NewNotification(app, notificationService, jwtMidd)
	api.NewTopUpApi(app, topUpService, jwtMidd)
	api.NewMidtrans(app, midtransService, topUpService)

	sse.NewNotification(app, jwtMidd, hub)
	log.Fatal(app.Listen(cnf.Server.Host + ":" + cnf.Server.Port))

}
