package dto

type GetClientReq struct {
	Id int64 `validate:"required"`
}

type GetClientRes struct {
	Id        int64
	GivenName string
	Surname   string
	Email     string
}
