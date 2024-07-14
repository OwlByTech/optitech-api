package dto

type UpdateDirectoryTreeReq struct {
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name"`
	InstitutionID int32  `json:"institutionId"`
}

type UpdateDirectoryTreeRes struct {
	DirectoryId   int64  `json:"directoryId"`
	InstitutionID int32  `json:"institutionId"`
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name"`
}
