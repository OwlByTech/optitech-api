package dto

type GetPermissionReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetPermissionRes struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
