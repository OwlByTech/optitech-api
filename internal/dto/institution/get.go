package dto

type GetInstitutionReq struct {
	InstitutionID int32 `validate:"required"`
}

type GetInstitutionRes struct {
	InstitutionID   int32
	InstitutionName string
	Logo            string
	Description     string
	Services        []string
}
