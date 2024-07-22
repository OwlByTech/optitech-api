package dto

type GetDirectoryRoleReq struct {
	UserId      int64 `json:"userId"`
	DirectoryId int64 `json:"directoryId"`
}

type GetDirectoryRoleRes struct {
	DirectoryId int64  `json:"directoryId"`
	UserId      int64  `json:"userId"`
	Status      string `json:"status"`
}
