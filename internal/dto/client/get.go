package dto

type GetClientReq struct {
	Id int64 `validate:"required"`
}

type GetClientRes struct {
	ClientID  int32
	GivenName string
	Surname   string
	Email     string
}
