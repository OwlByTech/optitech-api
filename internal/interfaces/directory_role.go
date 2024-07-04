package interfaces

import (
	models "optitech/internal/sqlc"
)

type IDirectoryRoleService interface {
	Create(req *[]models.CreateDirectoryRoleParams) error
}

type IDirectoryRoleRepository interface {
	CreateDirectoryRole(arg *[]models.CreateDirectoryRoleParams) error
}

type IDirectoryRoleHandler interface {
}
