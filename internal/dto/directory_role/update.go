package dto

type UpdateDirectoryRoleReq struct {
	DirectoryId int64  `json:"directoryId"`
	UserId      int64  `json:"userId"`
	Status      string `json:"status" validate:"required,oneof='r' 'w' 'x' 'rw' 'rx' 'wx' 'rwx'"`
}

type UpdateDirectoryRoleRes struct {
	DirectoryId int64 `json:"directoryId"`
	UserId      int64 `json:"userId"`
}
