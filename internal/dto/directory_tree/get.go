package dto

type GetDirectoryTreeReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDirectoryTreeRes struct {
	Id       int32  `json:"id"`
	ParentID int32  `json:"directoryId"`
	Name     string `jston:name`
}
