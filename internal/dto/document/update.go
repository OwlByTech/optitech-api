package dto

type UpdateDocumentReq struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name"`
	FileRute string `json:"fileRute"`
}

type UpdateDocumentRes struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	FileRute string `json:"fileRute"`
}
