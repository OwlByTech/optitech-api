package dto

type CreateFormatReq struct {
	AsesorId    int64    `json:"asesor_id" validate:"required"`
	FormatName  string   `json:"format_name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Items       []string `json:"items" validate:"required"`
	Extension   string   `json:"extension" validate:"required"`
	Version     string   `json:"version" validate:"required"`
}

type CreateFormatRes struct {
	Id          int64
	AsesorId    string
	FormatName  string
	Description string
	Items       []string
	Extension   string
	Version     string
}
