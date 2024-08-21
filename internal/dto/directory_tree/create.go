package dto

type CreateDirectoryTreeReq struct {
	ParentID      int64  `json:"parentId"`
	Name          string `json:"name" validate:"required"`
	InstitutionID int32  `json:"institutionId"`
	AsesorID      int32  `json:"asesorId"`
}

type CreateDirectoryTreeRes struct {
	DirectoryId   int64            `json:"directoryId"`
	InstitutionID int32            `json:"institutionId"`
	ParentID      int64            `json:"parentId"`
	Name          string           `json:"name"`
	Directories   []*DirectoryTree `json:"directories"`
}

type DirectoryTree struct {
	ID            int64            `json:"id"`
	ParentID      int64            `json:"parentId"`
	InstitutionID int32            `json:"institutionId"`
	AsesorID      int32            `json:"asesorId"`
	Name          string           `json:"name"`
	Directories   []*DirectoryTree `json:"directories"`
}
