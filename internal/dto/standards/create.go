package dto

type CreateStandardsReq struct {
	ServiceId  int32  `json:"service_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Complexity string `json:"complexity" validate:"required"`
	Modality   string `json:"modality" validate:"required"`
	Article    string `json:"article" validate:"required"`
	Section    string `json:"section" validate:"required"`
	Paragraph  string `json:"paragraph" validate:"required"`
	Criteria   string `json:"criteria" validate:"required"`
	Comply     bool   `json:"comply" validate:"required"`
	Applys     bool   `json:"applys" validate:"required"`
}

type CreateStandardsRes struct {
	Id         int64
	ServiceId  int32
	Name       string
	Complexity string
	Modality   string
	Article    string
	Section    string
	Paragraph  string
	Criteria   string
	Comply     bool
	Applys     bool
}
