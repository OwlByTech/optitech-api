package repository

import (
	"context"
	dto "optitech/internal/dto/notification"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryNotification struct {
	notificationRepository *sq.Queries
}

func NewrepositoryNotification(q *sq.Queries) interfaces.INotificationRepositoy {
	return &repositoryNotification{
		notificationRepository: q,
	}
}

func (r *repositoryNotification) CeateNotification(arg *sq.CreateNoficationParams) (*dto.CreateNotificationRes, error) {
	ctx := context.Background()

	res, err := r.notificationRepository.CreateNofication(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateNotificationRes{
		From:    dto.FromNotification(res.From),
		To:      dto.ToNotification(res.To),
		FromID:  res.FromID,
		ToID:    res.ToID,
		Message: res.Message,
		Title:   res.Title,
		Payload: res.Payload,
		Type:    dto.TypeNotification(res.Type.TypeNotification),
	}, nil
}

func (r *repositoryNotification) GetNotification(req *sq.GetNotificationParams) (*dto.GetNotificationRes, error) {
	ctx := context.Background()

	repoRes, err := r.notificationRepository.GetNotification(ctx, *req)

	if err != nil {
		return nil, err
	}

	return &dto.GetNotificationRes{
		ID:      repoRes.NotificationID,
		From:    dto.FromNotification(repoRes.From),
		To:      dto.ToNotification(repoRes.To),
		FromID:  repoRes.FromID,
		ToID:    repoRes.ToID,
		Message: repoRes.Message,
		Title:   repoRes.Title,
		Payload: repoRes.Payload,
		Type:    dto.TypeNotification(repoRes.Type.TypeNotification),
	}, nil
}
