package interfaces

import (
	p "optitech/internal/dto/permission"
	dto "optitech/internal/dto/role_permission"
	models "optitech/internal/sqlc"
)

type IRolePermissionService interface {
	GetByRoleId(roleId int32) (*[]p.GetPermissionRes, error)
	List() (*[]dto.GetRolePermissionRes, error)
}

type IRolePermissionRepository interface {
	GetRolePermission(id int64) (*dto.GetRolePermissionRes, error)
	ListPermissionByRoleId(roleId int32) (*[]models.ListPermissionByRoleIdRow, error)
	CreateRolePermission(arg *models.CreateRolePermissionParams) (*dto.CreateRolePermissionRes, error)
	UpdateRolePermission(arg *models.UpdateRolePermissionByIdParams) error
	ListRolePermissions() (*[]dto.GetRolePermissionRes, error)
	DeleteRolePermission(arg *models.DeleteRolePermissionByIdParams) error
}

type IRolePermissionHandler interface{}
