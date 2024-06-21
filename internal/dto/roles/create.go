package dto

type CreateRoleReq struct {
	RoleName    string `json:"role_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateRoleRes struct {
	Id          int64
	RoleName    string
	Description string
}
