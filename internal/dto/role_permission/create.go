package dto

type CreateRolePermissionReq struct {
	RoleId       string `json:"role_id" validate:"required"`
	PermissionId string `json:"permission_id" validate:"required"`
}

type CreateRolePermissionRes struct {
	Id           int64
	RoleId       int64
	PermissionId string
}
