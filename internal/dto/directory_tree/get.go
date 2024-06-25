package dto

type GetDirectoryTreeReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDirectoryTreeRes struct {
	Id          int64 `json:"id"`
	DirectoryId int64 `json:"directoryId"`
	RoleId      int64 `json:"roleId"`
}
