package dto

import "mime/multipart"

type StatusClient string

const (
	StatusClientActive   StatusClient = "activo"
	StatusClientInactive StatusClient = "inactivo"
)

type UpdateClientReq struct {
	ClientId  int32  `json:"id"`
	GivenName string `json:"givenName"`
	Surname   string `json:"surname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UpdateClientRes struct {
	GivenName string `json:"givenName" validate:"required"`
	Surname   string `json:"surname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}
type UpdateClientStatusReq struct {
	ClientId      int32        `json:"clientId" validate:"required"`
	Status        StatusClient `json:"status" validate:"required"`
	AsesorID      int32        `json:"asesorId" validate:"required"`
	InstitutionId int32        `json:"institutionId" validate:"required"`
}

type UpdateClientPhotoReq struct {
	ClientId int32
	Photo    *multipart.FileHeader
}
