package dto

type CreateClientReq struct {
	InstitutionName string   `json:"institutionName" validate:"required,institutionName"`
	Description     string   `json:"description" validate:"required"`
	Services        []string `json:"services" validate:"required"`
}

type CreateClientRes struct {
	Id              int64
	InstitutionName string
	Description     string
	Services        []string
}
