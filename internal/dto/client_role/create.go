package dto

type CreateClientRoleReq struct {
	ClientId int64 `json:"clientId" validate:"required"`
	RoleId   int64 `json:"roleId" validate:"required"`
}

type CreateClientRoleRes struct {
	Id       int64 `json:"id"`
	ClientId int32 `json:"clientId"`
	RoleId   int32 `json:"roleId"`
}
