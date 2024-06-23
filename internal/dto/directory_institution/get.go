package dto

type GetDirectoryInstitutionReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDirectoryInstitutionRes struct {
	Id            int64 `json:"id"`
	InstitutionId int64 `json:"institutionId"`
	DirectoryId   int64 `json:"directoryId"`
}
