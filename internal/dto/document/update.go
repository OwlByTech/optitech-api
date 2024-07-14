package dto

type UpdateDocumentReq struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" required`
}

type UpdateDocumentRes struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
