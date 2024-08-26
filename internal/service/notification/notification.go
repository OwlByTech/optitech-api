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

func (s *serviceNotification) Create(req *dto.CreateNotificationReq) (*dto.CreateNotificationRes, error) {
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
	return s.notificationRepository.GetNotification(req.ID)
}

func (s *serviceNotification) List() (*[]dto.GetNotificationRes, error) {
	repoRes, err := s.notificationRepository.ListNotifications()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceNotification) Update(req *dto.UpdateNotificationVisualizedReq) (bool, error) {
	notification, err := s.Get(dto.GetNotificationReq{ID: req.NotificationID})

	if err != nil {
		return false, err
	}

	visualized := pgtype.Bool{
		Bool:  notification.Visualized,
		Valid: true,
	}

	repoReq := &sq.UpdateNotificationVisualizedParams{
		NotificationID: req.NotificationID,
		Visualized:     visualized,
	}

	if req.Visualized != false {
		visualized = pgtype.Bool{
			Bool:  req.Visualized,
			Valid: true,
		}
	}

	err = s.notificationRepository.UpdateNotificationVisualized(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}
