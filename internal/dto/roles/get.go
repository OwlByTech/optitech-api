package dto

type GetRoleReq struct {
	Id int64 `validate:"required"`
}

type GetRoleRes struct {
	Id             int64
	PermissionType string
}
