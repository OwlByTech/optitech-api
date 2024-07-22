package dto

type CreateDirectoryRoleReq struct {
	DirectoryId int64  `json:"directoryId" validate:"required"`
	UserId      int64  `json:"userId" validate:"required"`
	Status      string `json:"status" validate:"required,oneof='r' 'w' 'x' 'rw' 'rx' 'wx' 'rwx'"`
}

type CreateDirectoryRoleRes struct {
	DirectoryId int64  `json:"directoryId"`
	UserId      int64  `json:"userId"`
	Status      string `json:"status"`
}
