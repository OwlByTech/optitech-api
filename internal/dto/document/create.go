package dto

import "mime/multipart"

type CreateDocumentReq struct {
	DirectoryId int32 `json:"directoryId" validate:"required"`
	FormatId    int32 `json:"formatId"`
	File        *multipart.FileHeader
	Status      string `json:"status" validate:"required"`
}

type CreateDocumentRes struct {
	Id          int64  `json:"id"`
	DirectoryId int32  `json:"directoryId"`
	Name        string `json:"name"`
	FormatId    int32  `json:"formatId"`
	FileRute    string `json:"fileRute"`
	Status      string `json:"status"`
}
