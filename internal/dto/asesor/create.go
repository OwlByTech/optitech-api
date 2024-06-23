package dto

type CreateAsesorReq struct {
	ClientId int32 `json:"client_id" validate:"required"`
	Photo    string
	About    string `json:"about" validate:"required"`
}

type CreateAsesorRes struct {
	Id    int32
	Photo string
	About string
}
