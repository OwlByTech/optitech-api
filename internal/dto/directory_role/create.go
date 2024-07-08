package dto

type CreateDirectoryRoleReq struct {
	DirectoryId int64  `json:"directoryId" validate:"required"`
	UserId      int64  `json:"roleId" validate:"required"`
	Status      string `json:"status" validate:"required,oneof='r' 'w' 'x' 'rw' 'rx' 'wx' 'rwx'"`
}

type CreateDirectoryRoleRes struct {
	DirectoryId int64  `json:"directoryId"`
	RoleId      int64  `json:"roleId"`
	Status      string `json:"status"`
}
