package dto

type CreateDirectoryTreeReq struct {
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name" validate:"required"`
	InstitutionID int64  `json:"institutionId":"required"`
}

type CreateDirectoryTreeRes struct {
	DirectoryId   int64  `json:"directoryId"`
	InstitutionID int64  `json:"institutionId"`
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name`
}
