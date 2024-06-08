package dto

type CreateServiceReq struct {
	ServiceName string `json:"service_name" validate:"required"`
}

type CreateServiceRes struct {
	Id          int64
	ServiceName string
}
