package dto

type UpdateServiceReq struct {
	Id          int64
	ServiceName string `json:"service_name" validate:"required"`
}
