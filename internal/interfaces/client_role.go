package interfaces

import (
	dto "optitech/internal/dto/client_role"
	r "optitech/internal/dto/roles"
	models "optitech/internal/sqlc"
)

type IClientRoleService interface {
	Create(arg *models.CreateClientRoleParams) (*dto.CreateClientRoleRes, error)
	Delete(arg *models.DeleteClientRoleByIdParams) error
	List() (*[]dto.GetClientRoleRes, error)
	ListByClientId(clientId int32) (*[]r.GetRoleRes, error)
}

type IClientRoleRepository interface {
	GetClientRole(id int64) (*dto.GetClientRoleRes, error)
	ListByClientId(id int32) (*[]models.GetClientRoleByClientIdRow, error)
	CreateClientRole(arg *models.CreateClientRoleParams) (*dto.CreateClientRoleRes, error)
	UpdateClientRole(arg *models.UpdateClientRoleByIdParams) error
	ListClientRoles() (*[]dto.GetClientRoleRes, error)
	DeleteClientRole(arg *models.DeleteClientRoleByIdParams) error
}

type IClientRoleHandler interface{}
