package dto

type GetAsesorReq struct {
	Id int32 `validate:"required"`
}

type GetAsesorRes struct {
	Id    int32  `json:"id"`
	About string `json:"about"`
}
