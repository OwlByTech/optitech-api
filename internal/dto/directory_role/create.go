package dto

type CreateDirectoryRoleReq struct {
	DirectoryId int64 `json:"directoryId" validate:"required"`
	RoleId      int64 `json:"roleId" validate:"required"`
}

type CreateDirectoryRoleRes struct {
	Id          int64 `json:"id"`
	DirectoryId int64 `json:"directoryId"`
	RoleId      int64 `json:"roleId"`
}
