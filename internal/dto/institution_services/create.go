package dto

type CreateInstitutionServicesReq struct {
	InstitutionID int64 `json:"institution_id" validate:"required"`
	ServiceId     int64 `json:"services_id" validate:"required"`
}

type CreateInstitutionServicesRes struct {
	Id            int64
	InstitutionID int64
	ServiceId     int64
}
