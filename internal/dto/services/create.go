package dto

type CreateServiceReq struct {
	Name string `json:"name" validate:"required"`
}

type CreateServiceRes struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
