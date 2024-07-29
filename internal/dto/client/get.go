package dto

import dto "optitech/internal/dto/roles"

type GetClientReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetClientRes struct {
	Id        int32        `json:"id"`
	GivenName string       `json:"givenName"`
	Photo     string       `json:"photo" `
	Surname   string       `json:"surname"`
	Status    StatusClient `json:"status"`
	Password  string
	Email     string           `json:"email"`
	Role      []dto.GetRoleRes `json:"roles"`
}
