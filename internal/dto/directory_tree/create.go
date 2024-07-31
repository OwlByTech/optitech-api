package dto

type CreateDirectoryTreeReq struct {
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name" validate:"required"`
	InstitutionID int32  `json:"institutionId"`
	AsesorID      int32  `json:"asesorId"`
}

type CreateDirectoryTreeRes struct {
	DirectoryId   int64  `json:"directoryId"`
	InstitutionID int32  `json:"institutionId"`
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name"`
}
