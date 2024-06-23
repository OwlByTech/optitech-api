package dto

type GetServiceReq struct {
	Id int32 `validate:"required"`
}

type GetServiceRes struct {
	Id   int32
	Name string
}
