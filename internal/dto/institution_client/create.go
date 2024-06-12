package dto

type CreateInstitutionClientReq struct {
	ClientId      string `json:"client_id" validate:"required"`
	InstitutionId string `json:"institution_id" validate:"required"`
}

type CreateInstitutionClientRes struct {
	Id            int64
	ClientId      int64
	InstitutionId int64
}
