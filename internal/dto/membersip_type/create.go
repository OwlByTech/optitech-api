package dto

type CreateMemberTypeshipReq struct {
	MembershipName string `json:"membership_name" validate:"required,membership_name"`
	Users          int64  `json:"users" validate:"required"`
}

type CreateMemberTypeshipRes struct {
	Id             int64
	MembershipName string
	Users          int64
}
