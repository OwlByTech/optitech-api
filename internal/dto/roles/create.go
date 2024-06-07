package dto

type CreateRoleReq struct {
	RoleName        string `json:"role_name" validate:"required"`
	RoleDescription string `json:"description" validate:"required"`
}

type CreateRoleRes struct {
	Id              int64
	RoleName        string
	RoleDescription string
}
