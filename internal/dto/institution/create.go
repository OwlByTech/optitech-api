package dto

type CreateInstitutionReq struct {
	InstitutionName string   `json:"institutionName" validate:"required"`
	Description     string   `json:"description" validate:"required"`
	Services        []string `json:"services" validate:"required"`
}

type CreateInstitutionRes struct {
	Id              int64    `json:"id"`
	InstitutionName string   `json:"institutionName"`
	Description     string   `json:"description"`
	Services        []string `json:"services"`
}
