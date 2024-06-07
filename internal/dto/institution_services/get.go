package dto

type GetInstitutionServicesReq struct {
	Id int64 `validate:"required"`
}

type GetInstitutionServicesRes struct {
	Id            int64
	InstitutionID int64
	ServicesId    int64
}
