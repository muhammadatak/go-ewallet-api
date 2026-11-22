package domain

import "errors"

var (
	ErrNotFound            = errors.New("Data is not found")
	ErrAuth                = errors.New("Wrong email or password")
	ErrInvalidToken        = errors.New("Invalid token")
	ErrUsernameExists      = errors.New("Username  already exists")
	ErrInsufficientBalance = errors.New("Insufficient balance")
	ErrInvalidOTP          = errors.New("Invalid OTP")
	ErrEmailNotVerified    = errors.New("Email is not verified")
)
