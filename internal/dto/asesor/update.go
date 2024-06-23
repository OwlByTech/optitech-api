package dto

type UpdateAsesorReq struct {
	AsesorID int32
	Photo    string `json:"photo" validate:"required"`
	About    string `json:"about" validate:"required"`
}

type UpdateAsesorRes struct {
	Id    int32
	Photo string
	About string
}
