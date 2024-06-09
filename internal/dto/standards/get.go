package dto

type GetStandardsReq struct {
	Id int64 `validate:"required"`
}

type GetStandardsRes struct {
	Id         int64
	ServiceId  int32
	Standard   string
	Complexity string
	Modality   string
	Article    string
	Section    string
	Paragraph  string
	Criteria   string
	Comply     bool
	Applys     bool
}
