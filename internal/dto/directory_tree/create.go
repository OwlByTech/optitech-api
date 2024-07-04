package dto

type CreateDirectoryTreeReq struct {
	ParentID int64  `json:"parentId"`
	Name     string `json:"name" validate:"required"`
}

type CreateDirectoryTreeRes struct {
	Id          int64  `json:"id"`
	DirectoryId int64  `json:"directoryId"`
	ParentID    int64  `json:"parentId"`
	Name        string `json:"name`
}
