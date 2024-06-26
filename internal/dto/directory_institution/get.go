package dto

type GetDirectoryInstitutionReq struct {
	Id int32 `json:"id" validate:"required"`
}

type GetDirectoryInstitutionRes struct {
	Id            int32 `json:"id"`
	InstitutionId int32 `json:"institutionId"`
	DirectoryId   int32 `json:"directoryId"`
}
