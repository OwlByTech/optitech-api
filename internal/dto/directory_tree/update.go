package dto

type UpdateDirectoryTreeReq struct {
	DirectoryId   int64  `json:"directoryId"`
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name"`
	InstitutionID int32  `json:"institutionId"`
	AsesorID      int32  `json:"asesorId"`
}

type UpdateDirectoryTreeRes struct {
	DirectoryId   int64  `json:"directoryId"`
	InstitutionID int32  `json:"institutionId"`
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name"`
	AsesorID      int32  `json:"asesorId"`
}
