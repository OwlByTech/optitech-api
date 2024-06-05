package dto

type MailingReq struct {
	Emails   []string `json:"emails" validate:"required,min=1,dive,email"`
	Password string   `json:"password" binding:"required"`
	Subject  string   `json:"subject" binding:"required"`
}
