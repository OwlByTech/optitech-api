package dto

type MailingReq struct {
	Email           string `json:"email" validate:"required,email"`
	PasswordMessage string `json:"password" validate:"required,min=6"`
	Subject         string `json:"subject" validate:"required,email"`
}

type CreateMailingRes struct {
}
