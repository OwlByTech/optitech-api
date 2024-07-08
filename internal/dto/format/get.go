package dto

type GetFormatReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetFormatRes struct {
	Id          int32    `json:"id"`
	AsesorId    int32    `json:"asesorId"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
	Extension   string   `json:"extension"`
	Version     string   `json:"version"`
}
