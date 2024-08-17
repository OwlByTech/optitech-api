package service

import (
	"fmt"
	ddto "optitech/internal/dto/document"
	dto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"optitech/internal/tools"
	"os"
	"time"

	ds "github.com/owlbytech/docu-stream-go"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceFormat struct {
	formatRepository interfaces.IFormatRepository
	documentService  interfaces.IDocumentService
}

func NewServiceFormat(f interfaces.IFormatRepository, serviceDocument interfaces.IDocumentService) interfaces.IFormatService {
	return &serviceFormat{
		formatRepository: f,
		documentService:  serviceDocument,
	}
}

func (s *serviceFormat) Get(req dto.GetFormatReq) (*dto.GetFormatRes, error) {
	return s.formatRepository.GetFormat(req.Id)
}

func (s *serviceFormat) Create(req *dto.CreateFormatReq) (*dto.CreateFormatRes, error) {
	repoReq := &sq.CreateFormatParams{
		AsesorID:    req.AsesorId,
		ServiceID:   pgtype.Int4{Int32: req.ServiceID, Valid: true},
		FormatName:  req.Name,
		Description: req.Description,
		Extension:   sq.Extensions(req.Extension),
		Version:     req.Version,
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if req.UpdateFormatID > 0 {
		repoReq.UpdatedFormatID = pgtype.Int4{Int32: req.UpdateFormatID, Valid: true}
	}

	r, err := s.formatRepository.CreateFormat(repoReq)
	if err != nil {
		return nil, err
	}

	fileByte, err := tools.FileToBytes(req.FormatFile)
	if err != nil {
		return nil, err
	}

	_, err = s.documentService.Create(&ddto.CreateDocumentByteReq{
		FormatId:    r.Id,
		DirectoryId: req.DirectoryId,
		File:        &fileByte,
		AsesorId:    r.AsesorId,
		Filename:    req.FormatFile.Filename,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceFormat) ListById(req *dto.ListFormatsReq) (*[]dto.GetFormatRes, error) {
	repoRes, err := s.formatRepository.ListById(&sq.ListFormatsByIdParams{
		Column1:  req.FormatsId,
		AsesorID: req.AsesorId,
	})
	if err != nil {
		return nil, err
	}
	return repoRes, nil
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

func (s *serviceFormat) ApplyWordFormat(req ds.WordApplyReq) ([]byte, error) {
	url := os.Getenv("DOCU_STREAM_IP")
	if url == "" {
		return nil, fmt.Errorf("WORD_CLIENT_URL not set in .env file")
	}
	c, err := ds.NewWordClient(&ds.ConnectOptions{Url: url})

	if err != nil {
		return nil, err
	}

	res, err := c.Apply(&req)

	if err != nil {
		return nil, err
	}

	return res.Docu, nil
}
