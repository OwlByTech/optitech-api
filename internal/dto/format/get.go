package dto

type GetFormatReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetFormatRes struct {
	Id          int64    `json:"id"`
	AsesorId    string   `json:"asesorId"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
	Extension   string   `json:"extension"`
	Version     string   `json:"version"`
}
