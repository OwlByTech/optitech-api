package dto

type GetRolePermissionReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetRolePermissionRes struct {
	Id           int64 `json:"id"`
	RoleId       int32 `json:"roleId"`
	PermissionId int32 `json:"permissionId"`
}
