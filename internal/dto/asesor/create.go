package dto

type CreateAsesorReq struct {
	ClientId int32  `json:"clientId" validate:"required"`
	About    string `json:"about" validate:"required"`
}

type CreateAsesorRes struct {
	Id    int32  `json:"id"`
	About string `json:"about"`
}
