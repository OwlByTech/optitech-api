package dto

type CreateDirectoryInstitutionReq struct {
	InstitutionID string `json:"institution_id" validate:"required"`
	DirectoryID   string `json:"directory_id" validate:"required"`
}

type CreateDirectoryInstitutionRes struct {
	Id            int64
	InstitutionID int64
	DirectoryID   int64
}
