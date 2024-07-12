package dto

type CreateDocumentClientReq struct {
	ClientId   int32  `json:"clientId" validate:"required"`
	DocumentId int32  `json:"documentId" validate:"required"`
	Action     string `json:"action" validate:"required"`
}

type CreateDocumentClientRes struct {
	Id         int64  `json:"id"`
	ClientId   int32  `json:"clientId"`
	DocumentId int32  `json:"documentId"`
	Action     string `json:"action"`
}
