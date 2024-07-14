package dto

import (
	"mime/multipart"
)

type CreateInstitutionReq struct {
	InstitutionName string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
	LogoFile        *multipart.FileHeader
	AsesorID        int32   `json:"asesorId"`
	Services        []int32 `json:"services" validate:"required"`
	Clients         []int32 `json:"clients"`
	DirectoryTree   []int64 `json:"directoryTree"`
}

type CreateInstitutionRes struct {
	InstitutionID int32 `json:"id"`
}
