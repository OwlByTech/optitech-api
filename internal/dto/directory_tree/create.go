package dto

type CreateDirectoryTreeReq struct {
	ParentID int32  `json:"parentId"`
	Name     string `json:"name" validate:"required"`
}

type CreateDirectoryTreeRes struct {
	Id          int32  `json:"id"`
	DirectoryId int32  `json:"directoryId"`
	ParentID    int32  `json:"parentId"`
	Name        string `json:"name"`
}
