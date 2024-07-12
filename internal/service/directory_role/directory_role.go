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

func (s *serviceDirectoryRole) Update(req *dto.UpdateDirectoryRoleReq) (bool, error) {
	directory, err := s.Get(dto.GetDirectoryRoleReq{UserId: req.UserId})

	if err != nil {
		return false, err
	}

	repoReq := &sq.UpdateDirectoryRoleParams{
		DirectoryID: pgtype.Int4{Int32: int32(directory.DirectoryId), Valid: true},
		UserID:      pgtype.Int4{Int32: int32(directory.UserId), Valid: true},
		Status:      sq.Permissions(req.Status),
	}

	if req.DirectoryId != 0 {
		repoReq.DirectoryID = pgtype.Int4{Int32: int32(req.DirectoryId), Valid: true}
	}

	if req.UserId != 0 {
		repoReq.DirectoryID = pgtype.Int4{Int32: int32(req.UserId), Valid: true}
	}

	if req.Status != "" {
		repoReq.Status = sq.Permissions(req.Status)
	}

	err = s.directoryRoleRepository.UpdateDirectoryRole(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceDirectoryRole) Delete(req dto.GetDirectoryRoleReq) (bool, error) {
	repoReq := &sq.DeleteDirectoryRoleByIdParams{
		UserID:      pgtype.Int4{Int32: int32(req.UserId), Valid: true},
		DirectoryID: pgtype.Int4{Int32: int32(req.DirectoryId), Valid: true},
		DeletedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.directoryRoleRepository.DeleteDirectoryRole(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}
