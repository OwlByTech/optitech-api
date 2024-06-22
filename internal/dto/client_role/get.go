package dto

import (
	c "optitech/internal/dto/client"
	r "optitech/internal/dto/roles"
)

type GetClientRoleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetClientRoleRes struct {
	Id       int64 `json:"id"`
	ClientId int32 `json:"clientId"`
	RoleId   int32 `json:"roleId"`
}

type GetClientRole struct {
	Client c.GetClientRes `json:"client"`
	Role   r.GetRoleRes   `json:"role"`
}
