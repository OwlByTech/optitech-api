package dto

type GetFormatReq struct {
	Id int64 `validate:"required"`
}

type GetFormatRes struct {
	Id          int64
	AsesorId    string
	Description string
	Items       []string
	Extension   string
	Version     string
}
