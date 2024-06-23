package dto

type UpdateInstitutionServicesReq struct {
	InstitutionId int32   `json:"institutionId" validate:"required"`
	Services      []int32 `json:"servicesId" validate:"required"`
}
