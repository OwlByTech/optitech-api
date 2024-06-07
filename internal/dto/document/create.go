package dto

type CreateDocumentClientReq struct {
	FormatId      int64  `json:"format_id" validate:"required"`
	InstitutionId int64  `json:"institution_id" validate:"required"`
	ClientId      int64  `json:"client_id" validate:"required"`
	FileRute      string `json:"file_rute" validate:"required"`
	Status        string `json:"status" validate:"required"`
}

type CreateDocumentClientRes struct {
	Id            int64
	FormatId      int64
	InstitutionId int64
	ClientId      int64
	FileRute      string
	Status        string
}
