package dto

type CreateRolePermissionReq struct {
	RoleId       int32 `json:"roleId" validate:"required"`
	PermissionId int32 `json:"permissionId" validate:"required"`
}

type CreateRolePermissionRes struct {
	Id           int64 `json:"id"`
	RoleId       int32 `json:"roleId"`
	PermissionId int32 `json:"permissionId"`
}
