package dto

type CreateDocumentClientClientReq struct {
	ClientId   int64  `json:"client_id" validate:"required"`
	DocumentId int64  `json:"document_id" validate:"required"`
	Action     string `json:"action" validate:"required"`
}

type CreateDocumentClientClientRes struct {
	Id         int64
	ClientId   int64
	DocumentId int64
	Action     string
}
