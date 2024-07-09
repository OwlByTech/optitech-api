package dto

type CreateFormatReq struct {
	AsesorId    int32    `json:"asesorId" validate:"required"`
	FormatName  string   `json:"formatName" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Items       []string `json:"items" validate:"required"`
	Extension   string   `json:"extension" validate:"required"`
	Version     string   `json:"version" validate:"required"`
}

type CreateFormatRes struct {
	Id          int32    `json:"id"`
	AsesorId    string   `json:"asesorId"`
	FormatName  string   `json:"formatName"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
	Extension   string   `json:"extension"`
	Version     string   `json:"version"`
}
