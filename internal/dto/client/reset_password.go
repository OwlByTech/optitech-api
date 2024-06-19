package dto

type ResetPasswordReq struct {
	Email string `json:"email" validate:"required,email"`
}
