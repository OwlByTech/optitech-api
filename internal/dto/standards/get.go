package dto

type GetStandardsReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetStandardsRes struct {
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
