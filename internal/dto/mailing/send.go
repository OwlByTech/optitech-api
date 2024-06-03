package dto

type MailingReq struct {
	Emails          []string `json:"emails" validate:"required,min=1,dive,email"`
	PasswordMessage string   `json:"passwordMessage" binding:"required"`
	Subject         string   `json:"subject" binding:"required"`
}
