package dto

type GetDocumentReq struct {
	Id            int64 `json:"id" validate:"required"`
	InstitutionId int32
	AsesorId      int32
}

type GetDocumentRes struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	DirectoryId int64  `json:"directoryId"`
	FormatId    int32  `json:"formatId"`
	FileRute    string `json:"fileRute"`
	Status      string `json:"status"`
}

type GetDocumentDownloadRes struct {
	AsesorId        int32  `json:"asesorId"`
	InstitutionId   int32  `json:"institutionId"`
	InstitutionName string `json:"name"`
	FileRute        string `json:"fileRute"`
	Filename        string `json:"filename"`
}
