package dto

type GetAsesorReq struct {
	Id int64 `validate:"required"`
}

type GetAsesorRes struct {
	Id       int64
	Username string
	Photo    string
	About    string
}
