package dto

type CreateAsesorReq struct {
	ClientId int32  `json:"clientId" validate:"required"`
	Username string `json:"username" validate:"required,username"`
	Photo    string `json:"photo" validate:"required"`
	About    string `json:"about" validate:"required"`
}

type CreateAsesorRes struct {
	Id       int32  `json:"id"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	About    string `json:"about"`
}
