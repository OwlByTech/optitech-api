package dto

type CreateDirectoryInstitutionReq struct {
	InstitutionID string `json:"institutionId" validate:"required"`
	DirectoryID   string `json:"directoryId" validate:"required"`
}

type CreateDirectoryInstitutionRes struct {
	Id            int64 `json:"id"`
	InstitutionID int64 `json:"institutionId"`
	DirectoryID   int64 `json:"directoryId"`
}
