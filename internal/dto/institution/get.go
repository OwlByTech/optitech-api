package dto

type GetInstitutionReq struct {
	Id int64 `validate:"required"`
}

type GetInstitutionRes struct {
	Id              int64
	InstitutionName string
	Logo            string
	Description     string
	Services        []string
}
