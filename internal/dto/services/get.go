package dto

type GetServiceReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetServiceRes struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
