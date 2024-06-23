package dto

type GetClientReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetClientRes struct {
	Id        int32  `json:"id"`
	GivenName string `json:"givenName"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}
