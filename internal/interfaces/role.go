package interfaces

import (
	dto "optitech/internal/dto/roles"
	models "optitech/internal/sqlc"
)

type IRoleService interface {
	List() (*[]dto.GetRoleRes, error)
}

type IRoleRepository interface {
	GetRole(id int32) (*dto.GetRoleRes, error)
	CreateRole(arg *models.CreateRoleParams) (*dto.CreateRoleRes, error)
	UpdateRole(arg *models.UpdateRoleByIdParams) error
	ListRoles() (*[]dto.GetRoleRes, error)
	DeleteRole(arg *models.DeleteRoleByIdParams) error
}

type IRoleHandler interface{}
