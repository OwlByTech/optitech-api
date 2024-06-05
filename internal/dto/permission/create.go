package dto

type CreatePermissionReq struct {
	PermissionType string `json:"permission_type" validate:"required"`
}

type CreatePermissionRes struct {
	Id             int64
	PermissionType string
}
