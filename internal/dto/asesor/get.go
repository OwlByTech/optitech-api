package dto

type GetAsesorReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetAsesorRes struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	About    string `json:"about"`
}
