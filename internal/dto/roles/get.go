package dto

type GetRoleReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetRoleRes struct {
	Id          int32  `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
