package dto

type CreateClientReq struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type CreateClientRes struct {
	Id        int64
	GivenName string
	Surname   string
	Email     string
}
