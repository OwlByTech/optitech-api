package interfaces

import (
	dto "optitech/internal/dto/client_role"
	models "optitech/internal/sqlc"
)

type IClientRoleService interface {
	List() (*[]dto.GetClientRoleRes, error)
}

type IClientRoleRepository interface {
	GetClientRole(id int64) (*dto.GetClientRoleRes, error)
	CreateClientRole(arg *models.CreateClientRoleParams) (*dto.CreateClientRoleRes, error)
	UpdateClientRole(arg *models.UpdateClientRoleByIdParams) error
	ListClientRoles() (*[]dto.GetClientRoleRes, error)
	DeleteClientRole(arg *models.DeleteClientRoleByIdParams) error
}

type IClientRoleHandler interface{}
