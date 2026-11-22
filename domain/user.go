package domain

import (
	"context"
	"database/sql"

	"github.com/muhammadatak/go-e-wallet/dto"
)

type User struct {
	Id              int64        `db:"id" goqu:"skipinsert,skipupdate"`
	Email           string       `d:"email"`
	Fullname        string       `db:"fullname"`
	Username        string       `db:"username"`
	Password        string       `db:"password"`
	Phone           string       `db:"phone"`
	EmailVerifiedAt sql.NullTime `db:"email_verified_at"`
}

type UserRepository interface {
	FindById(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Insert(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}

type UserService interface {
	Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error)
	ValidateOTP(ctx context.Context, req dto.ValidateOTPReq) error
}
