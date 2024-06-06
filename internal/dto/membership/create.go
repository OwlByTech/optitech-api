package dto

import "time"

type CreateMembershipReq struct {
	MembershipId int64     `json:"membership_id" validate:"required,membership_id"`
	FinishAt     time.Time `json:"finish_at" validate:"required"`
}

type CreateMembershipRes struct {
	Id           int64
	MembershipId int64
	FinishAt     time.Time
}
