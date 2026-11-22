package service

import (
	"context"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadatak/go-e-wallet/domain"
	"github.com/muhammadatak/go-e-wallet/dto"
	"github.com/muhammadatak/go-e-wallet/internal/config"
	"github.com/muhammadatak/go-e-wallet/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	cnf             *config.Config
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
	emailService    domain.EmailService
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository, cnf *config.Config, emailService domain.EmailService) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
		cnf:             cnf,
		emailService:    emailService,
	}
}
func (u *userService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	exists, err := u.userRepository.FindByUsername(ctx, req.Username)

	if err != nil {
		return dto.RegisterResponse{}, err
	}
	if exists.Id != 0 {
		return dto.RegisterResponse{}, domain.ErrUsernameExists
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	user := domain.User{
		Email:    req.Email,
		Fullname: req.Fullname,
		Username: req.Username,
		Password: string(hashPassword),
		Phone:    req.Phone,
	}
	err = u.userRepository.Insert(ctx, &user)

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	otpCode := util.GenerateOTPCode()
	referenceId := util.GenerateId()

	err = u.emailService.Send(req.Email, "OTP CODE", "otp anda"+otpCode)

	if err != nil {
		return dto.RegisterResponse{}, err
	}

	_ = u.cacheRepository.Set("otp:"+referenceId, []byte(otpCode))
	_ = u.cacheRepository.Set("user:"+referenceId, []byte(user.Username))
	return dto.RegisterResponse{
		ReferenceId: referenceId,
	}, nil
}

func (u *userService) ValidateOTP(ctx context.Context, req dto.ValidateOTPReq) error {
	code, err := u.cacheRepository.Get("otp:" + req.ReferenceId)
	username, err := u.cacheRepository.Get("user:" + req.ReferenceId)
	if err != nil {
		log.Println(err)
		return domain.ErrInvalidOTP
	}
	if string(code) != (req.OTPCode) {
		log.Println("calah")
		return domain.ErrInvalidOTP
	}

	user, err := u.userRepository.FindByUsername(ctx, string(username))
	if err != nil {
		return err
	}

	if user == (domain.User{}) {
		return domain.ErrNotFound
	}

	user.EmailVerifiedAt = sql.NullTime{Valid: true, Time: time.Now()}
	err = u.userRepository.Update(ctx, &user)

	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)

	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user == (domain.User{}) {
		return dto.AuthResponse{}, domain.ErrAuth
	}

	log.Println(req.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, domain.ErrAuth
	}
	if user.EmailVerifiedAt.Valid == false {
		return dto.AuthResponse{}, domain.ErrEmailNotVerified
	}
	token, err := issueToken(user.Id, u.cnf)

	if err != nil {
		return dto.AuthResponse{}, domain.ErrAuth
	}

	return dto.AuthResponse{
		Token: token,
	}, nil

}

func issueToken(userId int64, cnf *config.Config) (string, error) {
	claim := jwt.MapClaims{
		"sub": strconv.FormatInt(userId, 10),
		"exp": time.Now().Add(time.Duration(cnf.Jwt.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(cnf.Jwt.Key))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
