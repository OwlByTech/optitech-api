package dto

type GetClientRoleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetClientRoleRes struct {
	Id       int64 `json:"id"`
	ClientId int64 `json:"clientId"`
	RoleId   int64 `json:"roleId"`
}
