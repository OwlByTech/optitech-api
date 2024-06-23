package dto

type GetAsesorReq struct {
	Id int32 `validate:"required"`
}

type GetAsesorRes struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	About    string `json:"about"`
}
