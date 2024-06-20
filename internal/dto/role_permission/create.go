package dto

type CreateRolePermissionReq struct {
	RoleId       int32 `json:"role_id" validate:"required"`
	PermissionId int32 `json:"permission_id" validate:"required"`
}

type CreateRolePermissionRes struct {
	Id           int64
	RoleId       int32
	PermissionId int32
}
