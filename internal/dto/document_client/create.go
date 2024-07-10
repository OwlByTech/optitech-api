package dto

type CreateDocumentClientClientReq struct {
	ClientId   int32  `json:"clientId" validate:"required"`
	DocumentId int32  `json:"documentId" validate:"required"`
	Action     string `json:"action" validate:"required"`
}

type CreateDocumentClientClientRes struct {
	Id         int32  `json:"id"`
	ClientId   int32  `json:"clientId"`
	DocumentId int32  `json:"documentId"`
	Action     string `json:"action"`
}
