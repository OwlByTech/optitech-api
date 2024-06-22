package dto

type GetInstitutionClientReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetInstitutionClientRes struct {
	Id            int64 `json:"id"`
	ClientId      int64 `json:"clientId"`
	InstitutionId int64 `json:"institutionId"`
}
