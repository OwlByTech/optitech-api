package dto

type GetInstitutionClientReq struct {
	InstitutionId int32 `json:"institutionId"`
	ClientId      int32 `json:"clientId"`
}

type GetInstitutionClientRes struct {
	Id            int64 `json:"id"`
	ClientId      int64 `json:"clientId"`
	InstitutionId int64 `json:"institutionId"`
}
