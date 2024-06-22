package dto

type CreateInstitutionClientReq struct {
	ClientId      string `json:"clientId" validate:"required"`
	InstitutionId string `json:"institutionId" validate:"required"`
}

type CreateInstitutionClientRes struct {
	Id            int64 `json:"id"`
	ClientId      int64 `json:"clientId"`
	InstitutionId int64 `json:"institutionId"`
}
