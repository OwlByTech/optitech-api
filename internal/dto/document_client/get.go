package dto

type GetDocumentClientReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetDocumentClientRes struct {
	Id         int32  `json:"id"`
	ClientId   int32  `json:"clientId"`
	DocumentId int32  `json:"documentId"`
	Action     string `json:"action"`
}
