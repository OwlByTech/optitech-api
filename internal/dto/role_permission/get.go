package dto

type GetRolePermissionReq struct {
	Id int64 `validate:"required"`
}

type GetRolePermissionRes struct {
	Id           int64
	RoleId       int64
	PermissionId string
}
