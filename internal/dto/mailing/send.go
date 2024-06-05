package dto

type PasswordMailingReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" binding:"required"`
	Subject  string `json:"subject" binding:"required"`
}
