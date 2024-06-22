package interfaces

import (
	dto "optitech/internal/dto/role_permission"
	models "optitech/internal/sqlc"
)

type IRolePermissionService interface {
	List() (*[]dto.GetRolePermissionRes, error)
}

type IRolePermissionRepository interface {
	GetRolePermission(id int64) (*dto.GetRolePermissionRes, error)
	CreateRolePermission(arg *models.CreateRolePermissionParams) (*dto.CreateRolePermissionRes, error)
	UpdateRolePermission(arg *models.UpdateRolePermissionByIdParams) error
	ListRolePermissions() (*[]dto.GetRolePermissionRes, error)
	DeleteRolePermission(arg *models.DeleteRolePermissionByIdParams) error
}

type IRolePermissionHandler interface{}
