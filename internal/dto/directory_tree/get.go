package dto

type GetDirectoryRoleReq struct {
	Id int64 `validate:"required"`
}

type GetDirectoryRoleRes struct {
	Id          int64
	DirectoryID int64
	RoleId      int64
}
