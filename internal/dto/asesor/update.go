package dto

type UpdateAsesorReq struct {
	AsesorID int32
	About    string `json:"about" validate:"required"`
}
