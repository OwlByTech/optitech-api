package dto

type GetClientReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetClientRes struct {
	Id              int64    `json:"id"`
	InstitutionName string   `json:"institutionName"`
	Logo            string   `json:"logo"`
	Description     string   `json:"description"`
	Services        []string `json:"services"`
}
