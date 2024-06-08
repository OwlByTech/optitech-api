package dto

type GetInstitutionReq struct {
	InstitutionID int64 `validate:"required"`
}

type GetInstitutionRes struct {
	InstitutionID   int64
	InstitutionName string
	Logo            string
	Description     string
	Services        []string
}
