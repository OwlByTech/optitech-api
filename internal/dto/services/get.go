package dto

type GetServiceReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetServiceRes struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
