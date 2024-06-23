package dto

type GetDirectoryRoleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDirectoryRoleRes struct {
	Id          int64 `json:"id"`
	DirectoryId int64 `json:"directoryId"`
	RoleId      int64 `json:"roleId"`
}
