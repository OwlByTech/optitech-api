package dto

type CreateDirectoryTreeReq struct {
	ParentID int64  `json:"parentId" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type CreateDirectoryTreeRes struct {
	Id          int64 `json:"id"`
	DirectoryId int64 `json:"directoryId"`
}
