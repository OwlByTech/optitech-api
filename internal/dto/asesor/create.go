package dto

type CreateAsesorReq struct {
	ClientId int64  `json:"client_id" validate:"required"`
	Username string `json:"username" validate:"required,username"`
	Photo    string `json:"photo" validate:"required"`
	About    string `json:"about" validate:"required"`
}

type CreateAsesorRes struct {
	Id       int64
	Username string
	Photo    string
	About    string
}
