package dto

type CreateInstitutionClientReq struct {
	ClientId      string
	InstitutionId string
}

type CreateInstitutionClientRes struct {
	Id            int64
	ClientId      int64
	InstitutionId int64
}
