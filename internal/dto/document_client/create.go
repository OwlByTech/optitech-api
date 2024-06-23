package dto

type CreateDocumentClientClientReq struct {
	ClientId   int64  `json:"clientId" validate:"required"`
	DocumentId int64  `json:"documentId" validate:"required"`
	Action     string `json:"action" validate:"required"`
}

type CreateDocumentClientClientRes struct {
	Id         int64  `json:"id"`
	ClientId   int64  `json:"clientId"`
	DocumentId int64  `json:"documentId"`
	Action     string `json:"action"`
}
