package dto

import dto "optitech/internal/dto/document"

type GetDirectoryTreeReq struct {
	Id int64 `json:"id" validate:"required"`
}

type GetDirectoryTreeRes struct {
	Id       int32  `json:"id"`
	ParentID int32  `json:"parentId"`
	Name     string `json:"name"`
}

type GetDirectoryTreeByParentRes struct {
	Id        int32                 `json:"id"`
	ParentID  int32                 `json:"parentId"`
	Name      string                `json:"name"`
	Directory []GetDirectoryTreeRes `json:"directory"`
	Document  []dto.GetDocumentRes  `json:"document"`
}
