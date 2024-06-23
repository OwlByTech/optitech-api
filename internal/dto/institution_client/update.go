package dto

type UpdateInstitutionClientReq struct {
	InstitutionID int32   `json:"institutionId" validate:"required"`
	Clients       []int32 `json:"clients" validate:"required"`
}
