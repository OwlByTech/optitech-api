package dto

type GetClientReq struct {
	Id int64 `validate:"required"`
}

type GetClientRes struct {
	Id              int64
	InstitutionName string
	Logo            string
	Description     string
	Services        []string
}
