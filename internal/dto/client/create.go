package dto

type CreateClient struct {
	Id        int64
	GivenName string
	Surname   string
	Email     string
}

type CreateClientReq struct {
	GivenName string `json:"given_name" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

type CreateClientRes struct {
	Token string
}
