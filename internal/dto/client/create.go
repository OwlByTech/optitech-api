package dto

type CreateClient struct {
	Id        int32
	GivenName string
	Surname   string
	Email     string
}

type CreateClientReq struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	Role      int32  `json:"role" validate:"required"`
}

type CreateClientRes struct {
	Token string `json:"token"`
}
