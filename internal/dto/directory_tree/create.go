package dto

type CreateDirectoryRoleReq struct {
	DirectoryID int64 `json:"directory_id" validate:"required"`
	RoleId      int64 `json:"role_id" validate:"required"`
}

type CreateDirectoryRoleRes struct {
	Id          int64
	DirectoryID int64
	RoleId      int64
}
