package dto

type GetInstitutionClientReq struct {
	Id int64 `validate:"required"`
}

type GetInstitutionClientRes struct {
	Id            int64
	ClientId      int64
	InstitutionId int64
}
