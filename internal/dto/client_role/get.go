package dto

type GetClientRoleReq struct {
	Id int64 `validate:"required"`
}

type GetClientRoleRes struct {
	Id       int64
	ClientID int64
	RoleID   int64
}
