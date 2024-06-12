package dto

type CreatePermissionReq struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreatePermissionRes struct {
	Id          int64
	Name        string
	Code        string
	Description string
}
