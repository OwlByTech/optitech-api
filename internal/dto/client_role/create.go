package dto

type CreateClientRoleReq struct {
	ClientID int64 `json:"client_id" validate:"required"`
	RoleID   int64 `json:"role_id" validate:"required"`
}

type CreateClientRoleRes struct {
	Id       int64
	ClientID int64
	RoleID   int64
}
