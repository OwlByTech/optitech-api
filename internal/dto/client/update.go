package dto

type UpdateClientReq struct {
	ClientID  int32  `json:"client_id" validate:"required"`
	GivenName string `json:"given_name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
	Email     string `json:"email" validate:"required,email"`
}

type UpdateClientRes struct {
	GivenName string `json:"given_name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
