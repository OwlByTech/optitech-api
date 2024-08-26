package dto

type GetFormatReq struct {
	Id int32 `json:"id" validate:"required"`
}
type ListFormatsReq struct {
	FormatsId []int32 `json:"formatsId" validate:"required"`
	AsesorId  int32
}

type GetFormatRes struct {
	Id          int32    `json:"id"`
	AsesorId    int32    `json:"asesorId"`
	Description string   `json:"description"`
	Items       []string `json:"items"`
	Extension   string   `json:"extension"`
	Version     string   `json:"version"`
	Name        string   `json:"name"`
}
