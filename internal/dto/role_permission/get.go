package dto

import (
	p "optitech/internal/dto/permission"
	r "optitech/internal/dto/roles"
)

type GetRolePermissionReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetRolePermissionRes struct {
	Id           int32 `json:"id"`
	RoleId       int32 `json:"roleId"`
	PermissionId int32 `json:"permissionId"`
}

type GetRolePermission struct {
	Permission p.GetPermissionRes `json:"permission"`
	Role       r.GetRoleRes       `json:"role"`
}
