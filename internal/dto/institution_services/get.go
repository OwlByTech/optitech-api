package dto

type GetInstitutionServicesReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetInstitutionServicesRes struct {
	Id            int64 `json:"id"`
	InstitutionID int64 `json:"institutionId"`
	ServicesId    int64 `json:"servicesId"`
}
