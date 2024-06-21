package dto

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordTokenReq struct {
	Token string `json:"email" validate:"required,email"`
}
