package dto

type GetDocumentReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDocumentRes struct {
	Id          int64  `json:"id"`
	DirectoryId int32  `json:"directoryId"`
	FormatId    int32  `json:"formatId"`
	FileRute    string `json:"fileRute"`
	Status      string `json:"status"`
}
