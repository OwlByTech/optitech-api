package dto

type UpdateClientReq struct {
	ClientId  int64  `json:"clientId" validate:"required"`
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
	Email     string `json:"email" validate:"required,email"`
}

type UpdateClientRes struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
