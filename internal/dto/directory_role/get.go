package dto

type GetDirectoryRoleReq struct {
	DirectoryId int64 `json:"directoryId"`
	UserId      int64 `json:"roleId"`
}

type GetDirectoryRoleRes struct {
	DirectoryId int64  `json:"directoryId"`
	UserId      int64  `json:"roleId"`
	Status      string `json:"status"`
}
