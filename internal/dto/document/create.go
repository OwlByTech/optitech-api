package dto

type CreateDocumentClientReq struct {
	FormatId      int64  `json:"formatId" validate:"required"`
	InstitutionId int64  `json:"institutionId" validate:"required"`
	ClientId      int64  `json:"clientId" validate:"required"`
	FileRute      string `json:"fileRute" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

type CreateDocumentClientRes struct {
	Id            int64  `json:"id"`
	FormatId      int64  `json:"formatId"`
	InstitutionId int64  `json:"institutionId"`
	ClientId      int64  `json:"clientId"`
	FileRute      string `json:"fileRute"`
	Status        string `json:"status"`
}
