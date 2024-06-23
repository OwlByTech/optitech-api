package dto

type GetAsesorReq struct {
	Id int32 `validate:"required"`
}

type GetAsesorRes struct {
	Id    int32
	Photo string
	About string
}
