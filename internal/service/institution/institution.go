package service

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"io"
	dto "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"os"
	"time"
)

type service_institution struct {
	institutionRepository interfaces.IInstitutionRepository
}

func NewServiceInstitution(r interfaces.IInstitutionRepository) interfaces.IInstitutionService {
	return &service_institution{
		institutionRepository: r,
	}
}
func (s *service_institution) Get(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error) {
	repoRes, err := s.institutionRepository.GetInstitution(req.InstitutionID)

	if err != nil {
		return nil, err
	}

	return &dto.GetInstitutionRes{
		InstitutionID:   repoRes.InstitutionID,
		InstitutionName: repoRes.InstitutionName,
		Description:     repoRes.Description,
	}, nil
}

func (s *service_institution) List() ([]*dto.GetInstitutionRes, error) {
	repoRes, err := s.institutionRepository.ListInstitutions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (s *service_institution) Create(req *dto.CreateInstitutionReq) (*dto.CreateInstitutionRes, error) {
	repoReq := &sq.CreateInstitutionParams{
		InstitutionName: req.InstitutionName,
		Description:     req.Description,
		CreatedAt:       pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.LogoFile != nil {
		nameFile := req.InstitutionName + "_" + req.LogoFile.Filename
		multipart, err := req.LogoFile.Open()
		if err != nil {
			return nil, err
		}
		defer multipart.Close()
		savePath := fmt.Sprintf("./uploads/%s", nameFile)

		outFile, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer outFile.Close()
		if _, err = io.Copy(outFile, multipart); err != nil {
			return nil, err
		}
		repoReq.Logo = pgtype.Text{String: nameFile, Valid: true}
	}

	if req.AsesorID < 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}

	r, err := s.institutionRepository.CreateInstitution(repoReq)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *service_institution) Update(req *dto.UpdateInstitutionReq) (bool, error) {
	repoReq := &sq.UpdateInstitutionParams{
		InstitutionID: req.InstitutionID,
	}
	if req.AsesorID < 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}
	if req.InstitutionName != "" {
		repoReq.InstitutionName = req.InstitutionName
	}
	if req.Description != "" {
		repoReq.Description = req.Description
	}

	if req.LogoFile != nil {
		nameFile := req.InstitutionName + "_" + req.LogoFile.Filename
		multipart, err := req.LogoFile.Open()
		if err != nil {
			return false, err
		}
		defer multipart.Close()
		savePath := fmt.Sprintf("./uploads/%s", nameFile)

		outFile, err := os.Create(savePath)
		if err != nil {
			return false, err
		}
		defer outFile.Close()
		if _, err = io.Copy(outFile, multipart); err != nil {
			return false, err
		}
		repoReq.Logo = pgtype.Text{String: nameFile, Valid: true}
	}

	err := s.institutionRepository.UpdateInstitution(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service_institution) Delete(req dto.GetInstitutionReq) (bool, error) {
	repoReq := &sq.DeleteInstitutionParams{
		InstitutionID: req.InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err := s.institutionRepository.DeleteInstitution(repoReq)

	if err != nil {
		return false, err
	}

	return true, err
}
