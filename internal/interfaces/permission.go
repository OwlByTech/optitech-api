package interfaces

import (
	dto "optitech/internal/dto/permission"
	models "optitech/internal/sqlc"
)

type IPermissionService interface {
	List() (*[]dto.GetPermissionRes, error)
}

type IPermissionRepository interface {
	GetPermission(id int32) (*dto.GetPermissionRes, error)
	CreatePermission(arg *models.CreatePermissionParams) (*dto.CreatePermissionRes, error)
	UpdatePermission(arg *models.UpdatePermissionByIdParams) error
	ListPermissions() (*[]dto.GetPermissionRes, error)
	DeletePermission(arg *models.DeletePermissionByIdParams) error
}

type IPermissionHandler interface{}
