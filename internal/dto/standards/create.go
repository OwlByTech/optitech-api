package dto

type CreateStandardsReq struct {
	ServiceId  int32  `json:"serviceId" validate:"required"`
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
	Id         int64  `json:"id"`
	ServiceId  int32  `json:"serviceId"`
	Name       string `json:"name"`
	Complexity string `json:"complexity"`
	Modality   string `json:"modality"`
	Article    string `json:"article"`
	Section    string `json:"section"`
	Paragraph  string `json:"paragraph"`
	Criteria   string `json:"criteria"`
	Comply     bool   `json:"comply"`
	Applys     bool   `json:"applys"`
}
