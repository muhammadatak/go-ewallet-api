package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error when loading file configuration", err.Error())
	}

	expInt, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	if expInt == 0 {
		log.Fatal("jwt expired is not configured")
	}
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("jwt key is not configured")

	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USER"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TZ"),
		},
		Jwt: Jwt{
			Key: jwtKey,
			Exp: expInt,
		},
		Email: Email{
			Host:     os.Getenv("MAIL_HOST"),
			User:     os.Getenv("MAIL_USER"),
			Password: os.Getenv("MAIL_PASS"),
			Port:     os.Getenv("MAIL_PORT"),
		},
		Redis: Redis{
			Addr: os.Getenv("REDIS_ADDR"),
			Pass: os.Getenv("REDIS_PASS"),
		},
		Midtrans: Midtrans{
			Key:    os.Getenv("MIDTRANS_KEY"),
			IsProd: os.Getenv("MIDTRANS_ENV") == "production",
		},
	}
}
