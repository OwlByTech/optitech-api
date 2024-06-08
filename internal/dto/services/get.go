package dto

type GetServiceReq struct {
	ServiceID int64 `validate:"required"`
}

type GetServiceRes struct {
	ServiceID   int64
	ServiceName string
}
