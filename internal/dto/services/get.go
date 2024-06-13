package dto

type GetServiceReq struct {
	Id int32 `validate:"required"`
}

type GetServiceRes struct {
	ServiceID   int32
	ServiceName string
}
