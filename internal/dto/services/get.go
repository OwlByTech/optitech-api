package dto

type GetServiceReq struct {
	ServiceID int32 `validate:"required"`
}

type GetServiceRes struct {
	ServiceID   int32
	ServiceName string
}
