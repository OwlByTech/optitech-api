package dto

type GetMembershipTypeReq struct {
	Id int64 `validate:"required"`
}

type GetMembershipTypeRes struct {
	Id             int64
	MembershipName string
	Users          int64
}
