package dto

type CreateInstitutionServicesReq struct {
	InstitutionID int64 `json:"institutionId" validate:"required"`
	ServicesId    int64 `json:"servicesId" validate:"required"`
}

type CreateInstitutionServicesRes struct {
	Id            int64 `json:"id"`
	InstitutionID int64 `json:"institutionId"`
	ServicesId    int64 `json:"servicesId"`
}
