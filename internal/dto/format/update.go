package dto

type UpdateFormatReq struct {
	FormatID    int32  `json:"formatId"`
	FormatName  string `json:"formatName"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	Version     string `json:"version"`
	ServiceID   int32  `json:"serviceId"`
}

type UpdateFormatRes struct {
	FormatID    int32  `json:"formatId"`
	FormatName  string `json:"formatName"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	Version     string `json:"version"`
	ServiceID   int32  `json:"serviceId"`
}
