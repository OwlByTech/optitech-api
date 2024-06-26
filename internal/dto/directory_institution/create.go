package dto

type CreateDirectoryInstitutionReq struct {
	InstitutionID int32 `json:"institutionId" validate:"required"`
	DirectoryID   int32 `json:"directoryId" validate:"required"`
}

type CreateDirectoryInstitutionRes struct {
	Id            int32 `json:"id"`
	InstitutionID int32 `json:"institutionId"`
	DirectoryID   int32 `json:"directoryId"`
}
