package dto

type GetDocumentClientReq struct {
	Id int64 `validate:"required"`
}

type GetDocumentClientRes struct {
	Id         int64
	ClientId   int64
	DocumentId int64
	Action     string
}
