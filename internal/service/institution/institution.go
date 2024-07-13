package service

import (
	"fmt"
	"io"
	"log"
	cdto "optitech/internal/dto/client"
	dtdto "optitech/internal/dto/directory_tree"
	dto "optitech/internal/dto/institution"
	dto_services "optitech/internal/dto/institution_services"
	sdto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceInstitution struct {
	institutionRepository     interfaces.IInstitutionRepository
	serviceInstitutionService interfaces.IServiceInstitutionService
	clientInstitutionService  interfaces.IInstitutionClientService
	directoryTreeService      interfaces.IDirectoryService
	servicesService           interfaces.IService
}

func NewServiceInstitution(r interfaces.IInstitutionRepository, serviceInstitutionService interfaces.IServiceInstitutionService, serviceInstitutionClient interfaces.IInstitutionClientService, serviceDirectoryTree interfaces.IDirectoryService, services interfaces.IService) interfaces.IInstitutionService {
	return &serviceInstitution{
		institutionRepository:     r,
		serviceInstitutionService: serviceInstitutionService,
		clientInstitutionService:  serviceInstitutionClient,
		directoryTreeService:      serviceDirectoryTree,
		servicesService:           services,
	}
}

func (s *serviceInstitution) GetByClient(req cdto.GetClientReq) (*dto.GetInstitutionRes, error) {
	institutionId, err := s.institutionRepository.GetInstitutionByClient(req.Id)
	if err != nil {
		return nil, err
	}
	repoRes, err := s.Get(dto.GetInstitutionReq{Id: institutionId})
	if err != nil {
		return nil, err
	}
	return repoRes, err
}

func (s *serviceInstitution) Get(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error) {
	repoRes, err := s.institutionRepository.GetInstitution(req.Id)
	if err != nil {
		return nil, err
	}
	repoServices, _ := s.serviceInstitutionService.List(req.Id)
	repoClients, _ := s.clientInstitutionService.List(req.Id)
	repoRes.Clients = *repoClients
	repoRes.Services = *repoServices

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

	if req.AsesorID < 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}

	institutionID, err := s.institutionRepository.CreateInstitution(repoReq)

	if err != nil {
		return nil, err
	}

	services := make([]sq.CreateInstitutionServicesParams, len(req.Services))
	for i, ser := range req.Services {
		services[i] = sq.CreateInstitutionServicesParams{
			InstitutionID: institutionID,
			ServiceID:     ser,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
		}
	}

	if err := s.serviceInstitutionService.Create(&services); err != nil {
		return nil, err
	}

	clients := make([]sq.CreateInstitutionClientParams, len(req.Clients))
	for i, client := range req.Clients {
		clients[i] = sq.CreateInstitutionClientParams{
			InstitutionID: institutionID,
			ClientID:      client,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
		}
	}

	if err := s.clientInstitutionService.Create(&clients); err != nil {
		return nil, err
	}

	directoryTreeReq := &dtdto.CreateDirectoryTreeReq{
		InstitutionID: institutionID,
		Name:          repoReq.InstitutionName,
	}

	rootDirectoryID, err := s.directoryTreeService.Create(directoryTreeReq)
	if err != nil {
		return nil, err
	}

	for _, serviceID := range req.Services {

		getServiceReq := sdto.GetServiceReq{
			Id: serviceID,
		}

		log.Print(getServiceReq)
		log.Print(rootDirectoryID.DirectoryId)

		serviceName, err := s.servicesService.Get(getServiceReq)
		if err != nil {
			return nil, err
		}

		serviceDirectoryReq := &dtdto.CreateDirectoryTreeReq{
			ParentID:      rootDirectoryID.DirectoryId,
			Name:          serviceName.Name,
			InstitutionID: institutionID,
		}
		_, err = s.directoryTreeService.Create(serviceDirectoryReq)
		if err != nil {
			return nil, err
		}
	}

	return &dto.CreateInstitutionRes{InstitutionID: institutionID}, nil
}

func (s *serviceInstitution) UpdateAsesor(req *dto.UpdateAsesorInstitutionReq) (bool, error) {
	repoReq := &sq.UpdateAsesorInstitutionParams{
		InstitutionID: req.InstitutionID,
		AsesorID:      pgtype.Int4{Int32: req.AsesorID, Valid: true},
		UpdatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if err := s.institutionRepository.UpdateAsesorInstitution(repoReq); err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceInstitution) Update(req *dto.UpdateInstitutionReq) (bool, error) {
	institution, err := s.Get(dto.GetInstitutionReq{Id: req.InstitutionID})
	repoReq := &sq.UpdateInstitutionParams{
		InstitutionID:   req.InstitutionID,
		InstitutionName: institution.InstitutionName,
		Description:     institution.Description,
		UpdatedAt:       pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if req.InstitutionName != "" {
		repoReq.InstitutionName = req.InstitutionName
	}
	if req.Description != "" {
		repoReq.Description = req.Description
	}

	err = s.institutionRepository.UpdateInstitution(repoReq)

	if err != nil {
		return false, err
	}
	err = s.serviceInstitutionService.Update(&dto_services.UpdateInstitutionServicesReq{InstitutionId: req.InstitutionID, Services: req.Services})

	if err != nil {
		return false, err
	}

	return true, nil
}
func (s *serviceInstitution) UpdateLogo(req *dto.UpdateLogoReq) (bool, error) {
	repoReq := &sq.UpdateLogoInstitutionParams{
		InstitutionID: req.InstitutionID,
		UpdatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	if req.LogoFile != nil {
		nameFile := strconv.Itoa(int(req.InstitutionID)) + "_institution_" + req.LogoFile.Filename
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
		//		defer outFilef.Close()
		if _, err = io.Copy(outFile, multipart); err != nil {
			return false, err
		}
		repoReq.Logo = pgtype.Text{String: nameFile, Valid: true}
	}

	err := s.institutionRepository.UpdateLogoInstitution(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}
func (s *serviceInstitution) Delete(req dto.GetInstitutionReq) (bool, error) {
	repoReq := &sq.DeleteInstitutionParams{
		InstitutionID: req.Id,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.institutionRepository.DeleteInstitution(repoReq); err != nil {
		return false, err
	}
	if err := s.clientInstitutionService.DeleteByInstitution(repoReq.InstitutionID); err != nil {
		return false, err
	}
	if err := s.serviceInstitutionService.DeleteByInstitution(repoReq.InstitutionID); err != nil {
		return false, err
	}

	return true, nil
}
