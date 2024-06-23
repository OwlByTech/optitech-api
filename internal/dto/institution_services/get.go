package dto

type GetInstitutionServicesReq struct {
	Id            int64 `json:"id"`
	InstitutionId int32 `json:"institutionId"`
	ServiceId     int32 `json:"servicesId"`
}

type GetInstitutionServicesRes struct {
	Id int64 `json:"id" validate:"required"`
}
