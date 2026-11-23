package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3"`
}

type AuthResponse struct {
	Token  string `json:"token"`
	UserId int64  `json:"-"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=72"`
	Phone    string `json:"phone"`
}

type RegisterResponse struct {
	ReferenceId string `json:"reference_id"`
}

type ValidateOTPReq struct {
	ReferenceId string `json:"reference_id"`
	OTPCode     string `json:"otp_code"`
}
