package dto

import "mime/multipart"

type CreateDocumentReq struct {
	DirectoryId   int64  `json:"directoryId" validate:"required"`
	FormatId      int32  `json:"formatId"`
	Status        string `json:"status" validate:"required"`
	File          *multipart.FileHeader
	AsesorId      int32
	InstitutionId int32
}

type CreateDocumentRes struct {
	Id          int64  `json:"id"`
	DirectoryId int64  `json:"directoryId"`
	Name        string `json:"name"`
	FormatId    int32  `json:"formatId"`
	FileRute    string `json:"fileRute"`
	Status      string `json:"status"`
}

type CreateDocumentByteReq struct {
	DirectoryId   int64
	FormatId      int32
	File          *[]byte
	Filename      string
	Status        string
	AsesorId      int32
	InstitutionId int32
}

type CreateDocumentVersionReq struct {
	Id            int64 `json:"id"`
	File          *multipart.FileHeader
	AsesorId      int32
	InstitutionId int32
}

type CreateDocumentVersionByteReq struct {
	Id            int64
	Filename      string
	File          *[]byte
	Status        string
	AsesorId      int32
	InstitutionId int32
}
