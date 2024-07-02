package dto

type UpdateClientReq struct {
	ClientId  int32  `json:"id"`
	GivenName string `json:"givenName"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Token     string `json:"token" validate:"required"`
}

type UpdateClientRes struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
