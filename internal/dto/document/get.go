package dto

type GetDocumentReq struct {
	Id int64 `validate:"required"`
}

type GetDocumentRes struct {
	Id            int64
	FormatId      int64
	InstitutionId int64
	ClientId      int64
	FileRute      string
	Status        string
}
