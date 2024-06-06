package dto

import "time"

type GetMembershipReq struct {
	Id int64 `validate:"required"`
}

type GetMembershipRes struct {
	Id           int64
	MembershipId int64
	FinishAt     time.Time
}
