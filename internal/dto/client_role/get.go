package dto

type GetClientRoleReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetClientRoleRes struct {
	Id       int64 `json:"id"`
	ClientId int32 `json:"clientId"`
	RoleId   int32 `json:"roleId"`
}
