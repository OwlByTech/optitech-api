package dto

type CreateDirectoryTreeReq struct {
	ParentID int64  `json:"parentId"`
	Name     string `json:"name" validate:"required"`
}

type CreateDirectoryTreeRes struct {
	DirectoryId int64  `json:"directoryId"`
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name`
}
