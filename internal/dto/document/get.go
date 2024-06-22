package dto

type GetDocumentReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDocumentRes struct {
	Id            int64  `json:"id"`
	FormatId      int64  `json:"formatId"`
	InstitutionId int64  `json:"institutionId"`
	ClientId      int64  `json:"clientId"`
	FileRute      string `json:"fileRute"`
	Status        string `json:"status"`
}
