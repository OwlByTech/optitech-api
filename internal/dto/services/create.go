package dto

type CreateRoleReq struct {
	ServiceName string `json:"service_name" validate:"required"`
}

type CreateRoleRes struct {
	Id          int64
	ServiceName string
}
