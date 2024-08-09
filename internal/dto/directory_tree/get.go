package dto

import dto "optitech/internal/dto/document"

type GetDirectoryTreeReq struct {
	Id            int64 `json:"id" validate:"required"`
	InstitutionID int32 `json:"institutionId"`
}

type GetDirectoryTreeRes struct {
	Id            int64                  `json:"id"`
	ParentID      int64                  `json:"parentId"`
	InstitutionID int32                  `json:"institutionId"`
	Open          bool                   `json:"open"`
	Name          string                 `json:"name"`
	Directory     []*GetDirectoryTreeRes `json:"directory"`
	Document      *[]dto.GetDocumentRes  `json:"document"`
	AsesorID      int32                  `json:"asesorId"`
}
