package dto

type GetPermissionReq struct {
	Id int64 `validate:"required"`
}

type GetPermissionRes struct {
	Id             int64
	PermissionType string
}
