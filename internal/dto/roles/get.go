package dto

type GetRoleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetRoleRes struct {
	Id          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
