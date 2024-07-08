package dto

type UpdateDirectoryRoleReq struct {
	DirectoryId int64 `json:"directoryId"`
	UserId      int64 `json:"userId"`
}

type UpdateDirectoryRoleRes struct {
	DirectoryId int64 `json:"directoryId"`
	UserId      int64 `json:"userId"`
}
