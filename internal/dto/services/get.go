package dto

type GetServiceReq struct {
	Id int64 `validate:"required"`
}

type GetServiceRes struct {
	Id          int64
	ServiceName string
}
