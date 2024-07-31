package service

import (
	dto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceFormat struct {
	formatRepository interfaces.IFormatRepository
}

func NewServiceFormat(f interfaces.IFormatRepository) interfaces.IFormatService {
	return &serviceFormat{
		formatRepository: f,
	}
}

func (s *serviceFormat) Get(req dto.GetFormatReq) (*dto.GetFormatRes, error) {
	return s.formatRepository.GetFormat(req.Id)
}

func (s *serviceFormat) Create(req *dto.CreateFormatReq) (*dto.CreateFormatRes, error) {
	var updateFormatID pgtype.Int4

	if req.UpdateFormatID == 0 {
		updateFormatID.Valid = false
	} else {
		updateFormatID.Int32 = int32(req.UpdateFormatID)
		updateFormatID.Valid = true
	}

	repoReq := &sq.CreateFormatParams{
		UpdatedFormatID: pgtype.Int4{Int32: req.UpdateFormatID, Valid: false},
		AsesorID:        req.AsesorId,
		ServiceID:       pgtype.Int4{Int32: req.ServiceID, Valid: true},
		FormatName:      req.Name,
		Description:     req.Description,
		Extension:       sq.Extensions(req.Extension),
		Version:         req.Version,
		CreatedAt:       pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.formatRepository.CreateFormat(repoReq)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceFormat) List() (*[]dto.GetFormatRes, error) {
	repoRes, err := s.formatRepository.List()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceFormat) Delete(req dto.GetFormatReq) (bool, error) {
	repoReq := &sq.DeleteFormatByIdParams{
		FormatID:  req.Id,
		DeletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.formatRepository.DeleteFormat(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}

func (s *serviceFormat) Update(req *dto.UpdateFormatReq) (bool, error) {
	format, err := s.Get(dto.GetFormatReq{Id: req.FormatID})

	if err != nil {
		return false, err
	}

	repoReq := &sq.UpdateFormatByIdParams{
		FormatID:    req.FormatID,
		FormatName:  format.Name,
		Description: format.Description,
		Extension:   sq.Extensions(format.Extension),
		Version:     format.Version,
		UpdatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.Name != "" {
		repoReq.FormatName = req.Name
	}

	if req.Description != "" {
		repoReq.Description = req.Description
	}

	if req.Extension != "" {
		repoReq.Extension = sq.Extensions(req.Extension)
	}

	if req.Version != "" {
		repoReq.Version = req.Version
	}

	err = s.formatRepository.UpdateFormat(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}
