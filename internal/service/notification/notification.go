package service

import (
	dto "optitech/internal/dto/notification"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceNotification struct {
	notificationRepository interfaces.INotificationRepositoy
}

func NewServiceNotification(r interfaces.INotificationRepositoy) interfaces.INotificationService {
	return &serviceNotification{
		notificationRepository: r,
	}
}

func (s *serviceNotification) Create(req *dto.CreateNorificationReq) (*dto.CreateNotificationRes, error) {
	repoReq := &sq.CreateNoficationParams{
		From:      sq.FromNotification(req.From),
		To:        sq.ToNotification(req.To),
		FromID:    req.FromID,
		ToID:      req.FromID,
		Message:   req.Message,
		Title:     req.Title,
		Payload:   req.Payload,
		Type:      sq.NullTypeNotification{},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.notificationRepository.CeateNotification(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceNotification) Get(req dto.GetNotificationReq) (*dto.GetNotificationRes, error) {
	return s.notificationRepository.GetNotification(&sq.GetNotificationParams{
		NotificationID: req.ID,
		ToID:           int32(req.ID),
	})
}
