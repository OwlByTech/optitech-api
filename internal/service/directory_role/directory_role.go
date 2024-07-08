package service

import (
	dto "optitech/internal/dto/directory_role"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceDirectoryRole struct {
	directoryRoleRepository interfaces.IDirectoryRoleRepository
}

func NewServiceDirectoryRole(r interfaces.IDirectoryRoleRepository) interfaces.IDirectoryRoleService {
	return &serviceDirectoryRole{
		directoryRoleRepository: r,
	}
}

func (s *serviceDirectoryRole) Create(req *dto.CreateDirectoryRoleReq) (*dto.CreateDirectoryRoleRes, error) {

	repoReq := &sq.CreateDirectoryRoleParams{
		DirectoryID: pgtype.Int4{Int32: int32(req.DirectoryId), Valid: true},
		UserID:      pgtype.Int4{Int32: int32(req.UserId), Valid: true},
		Status:      sq.Permissions(req.Status),
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.directoryRoleRepository.CreateDirectoryRole(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceDirectoryRole) Get(req dto.GetDirectoryRoleReq) (*dto.GetDirectoryRoleRes, error) {
	return s.directoryRoleRepository.GetDirectoryRole(req.UserId)
}

func (s *serviceDirectoryRole) List() (*[]dto.GetDirectoryRoleRes, error) {
	repoRes, err := s.directoryRoleRepository.ListDirectoryRole()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}
