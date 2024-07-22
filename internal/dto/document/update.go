package dto

type UpdateDocumentReq struct {
	Id          int64  `json:"id" validate:"required"`
	Name        string `json:"name"`
	DirectoryID int64  `json:"directory_id"`
	FileRute    string `json:"fileRute"`
}

type UpdateDocumentRes struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	DirectoryID int64  `json:"directory_id"`
	FileRute    string `json:"fileRute"`
}
