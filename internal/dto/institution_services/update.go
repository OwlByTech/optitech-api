package dto

type UpdateInstitutionServicesReq struct {
	InstitutionID int32   `json:"institution_id" validate:"required"`
	Services      []int32 `json:"services_id" validate:"required"`
}
