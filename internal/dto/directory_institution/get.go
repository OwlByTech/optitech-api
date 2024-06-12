package dto

type GetDirectoryInstitutionReq struct {
	Id int64 `validate:"required"`
}

type GetDirectoryInstitutionRes struct {
	Id            int64
	InstitutionID int64
	DirectoryID   int64
}
