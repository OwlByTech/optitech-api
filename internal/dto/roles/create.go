package dto

type CreateRoleReq struct {
	RoleName    string `json:"roleName" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateRoleRes struct {
	Id          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
