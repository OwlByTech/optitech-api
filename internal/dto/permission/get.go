package dto

type GetPermissionReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetPermissionRes struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
