package dto

type CreatePermissionReq struct {
	PermissionName        string `json:"permission_name" validate:"required"`
	PermissionCode        string `json:"permission_code" validate:"required"`
	PermissionDescription string `json:"permission_description" validate:"required"`
}

type CreatePermissionRes struct {
	Id                    int64
	PermissionName        string
	PermissionCode        string
	PermissionDescription string
}
