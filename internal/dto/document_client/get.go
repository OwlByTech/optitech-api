package dto

type GetDocumentClientReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDocumentClientRes struct {
	Id         int64  `json:"id"`
	ClientId   int64  `json:"clientId"`
	DocumentId int64  `json:"documentId"`
	Action     string `json:"action"`
}
