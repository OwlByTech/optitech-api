package dto

type UpdateServiceReq struct {
	ServiceID   int32
	ServiceName string `json:"service_name" validate:"required"`
}
