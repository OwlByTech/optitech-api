package dto

type CreateFormatReq struct {
	UpdateFormatID int32  `json:updateFormatId`
	AsesorId       int32  `json:"asesorId" validate:"required"`
	ServiceID      int32  `json:"serviceId"`
	FormatName     string `json:"formatName" validate:"required"`
	Description    string `json:"description" validate:"required"`
	Extension      string `json:"extension" validate:"required"`
	Version        string `json:"version" validate:"required"`
}

type CreateFormatRes struct {
	Id          int32  `json:"id"`
	AsesorId    string `json:"asesorId"`
	FormatName  string `json:"formatName"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	Version     string `json:"version"`
}
