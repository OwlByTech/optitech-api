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

type serviceInstitution struct {
	institutionRepository     interfaces.IInstitutionRepository
	serviceInstitutionService interfaces.IServicesInstitutionService
	clientInstitutionService  interfaces.IInstitutionClientService
}

func NewServiceInstitution(r interfaces.IInstitutionRepository, serviceInstitutionService interfaces.IServicesInstitutionService, serviceInstitutionClient interfaces.IInstitutionClientService) interfaces.IInstitutionService {
	return &serviceInstitution{
		institutionRepository:     r,
		serviceInstitutionService: serviceInstitutionService,
		clientInstitutionService:  serviceInstitutionClient,
	}
}
func (s *serviceInstitution) Get(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error) {
	repoRes, err := s.institutionRepository.GetInstitution(req.InstitutionID)

	if err != nil {
		return nil, err
	}

	return repoRes, err
}

func (s *serviceInstitution) List() (*[]dto.GetInstitutionRes, error) {
	repoRes, err := s.institutionRepository.ListInstitutions()
	if err != nil {
		return nil, err
	}

	return repoRes, nil
}

func (s *serviceInstitution) Create(req *dto.CreateInstitutionReq) (*dto.CreateInstitutionRes, error) {
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
	services := make([]sq.CreateInstitutionServicesParams, len(req.Services))
	for i, ser := range req.Services {
		services[i] = sq.CreateInstitutionServicesParams{
			InstitutionID: r.InstitutionID,
			ServiceID:     ser,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
		}
	}

	if err := s.serviceInstitutionService.Create(&services); err != nil {
		return nil, err
	}
	if services, err := s.serviceInstitutionService.List(r.InstitutionID); err != nil {
		return nil, err
	} else {
		r.Services = *services
	}
	clients := make([]sq.CreateInstitutionClientParams, len(req.Clients))
	for i, client := range req.Clients {
		clients[i] = sq.CreateInstitutionClientParams{
			InstitutionID: r.InstitutionID,
			ClientID:      client,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
		}
	}

	if err := s.clientInstitutionService.Create(&clients); err != nil {
		return nil, err
	}
	if clients, err := s.clientInstitutionService.List(r.InstitutionID); err != nil {
		return nil, err
	} else {
		r.Clients = *clients
	}

	return r, nil
}

func (s *serviceInstitution) Update(req *dto.UpdateInstitutionReq) (bool, error) {
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

func (s *serviceInstitution) Delete(req dto.GetInstitutionReq) (bool, error) {
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
