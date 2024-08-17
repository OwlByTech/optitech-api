package service

import (
	"fmt"
	cdto "optitech/internal/dto/client"
	dtdto "optitech/internal/dto/directory_tree"
	ddto "optitech/internal/dto/document"
	dto "optitech/internal/dto/institution"
	dto_services "optitech/internal/dto/institution_services"
	sdto "optitech/internal/dto/services"
	"optitech/internal/interfaces"
	digitalOcean "optitech/internal/service/digital_ocean"
	sq "optitech/internal/sqlc"
	"optitech/internal/tools"
	"strconv"
	"time"

	docustream "github.com/owlbytech/docu-stream-go"

	ds "github.com/owlbytech/docu-stream-go"

	"github.com/jackc/pgx/v5/pgtype"
)

const asesorEnum = "Asesor"

type serviceInstitution struct {
	institutionRepository     interfaces.IInstitutionRepository
	serviceInstitutionService interfaces.IServiceInstitutionService
	clientInstitutionService  interfaces.IInstitutionClientService
	directoryTreeService      interfaces.IDirectoryService
	servicesService           interfaces.IService
	formatService             interfaces.IFormatService
	documentsService          interfaces.IDocumentService
}

func NewServiceInstitution(r interfaces.IInstitutionRepository, serviceInstitutionService interfaces.IServiceInstitutionService, serviceInstitutionClient interfaces.IInstitutionClientService, serviceDirectoryTree interfaces.IDirectoryService, services interfaces.IService, serviceFormat interfaces.IFormatService, serviceDocument interfaces.IDocumentService) interfaces.IInstitutionService {
	return &serviceInstitution{
		institutionRepository:     r,
		serviceInstitutionService: serviceInstitutionService,
		clientInstitutionService:  serviceInstitutionClient,
		directoryTreeService:      serviceDirectoryTree,
		servicesService:           services,
		formatService:             serviceFormat,
		documentsService:          serviceDocument,
	}
}

func (s *serviceInstitution) GetByClient(req cdto.GetClientReq) (int32, error) {
	return s.institutionRepository.GetInstitutionByClient(req.Id)
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
func (s *serviceInstitution) GetLogo(req dto.GetInstitutionReq) (string, error) {
	institution, err := s.institutionRepository.GetInstitutionLogo(req.Id)
	if err != nil {
		return "", err
	}
	route, err := digitalOcean.DownloadDocument(institution.Logo, institution.InstitutionName)
	if err != nil {
		return "", err
	}
	return route, nil
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
	institution, err := s.Get(dto.GetInstitutionReq{Id: req.InstitutionID})
	if err != nil {
		return false, err
	}
	repoReq := &sq.UpdateLogoInstitutionParams{
		InstitutionID: req.InstitutionID,
		UpdatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	logo, err := tools.FileToBytes(req.LogoFile)
	if err != nil {
		return false, err
	}

	if req.LogoFile != nil {
		name, err := digitalOcean.UploadDocument(logo, institution.InstitutionName, institution.InstitutionName)
		if err != nil {
			return false, err
		}

		repoReq.Logo = pgtype.Text{String: *name, Valid: true}

	}

	err = s.institutionRepository.UpdateLogoInstitution(repoReq)

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

func (s *serviceInstitution) CreateAllFormat(req *dto.GetInstitutionReq) (bool, error) {
	institution, err := s.Get(*req)
	if err != nil {
		return false, err
	}

	asesorId := institution.AsesorId
	if asesorId == 0 {
		return false, fmt.Errorf("Asesor not found")
	}

	directoryAsesor := dtdto.GetDirectoryTreeReq{
		AsesorID: asesorId,
	}

	directoryAsesorParent, err := s.directoryTreeService.GetIdByParent(&directoryAsesor)
	if err != nil {
		return false, err
	}

	directoryInstitution := dtdto.GetDirectoryTreeReq{
		InstitutionID: institution.Id,
	}

	directoryInstitutionParentId, err := s.directoryTreeService.GetIdByParent(&directoryInstitution)
	if err != nil {
		return false, err
	}

	directoryAsesor.Id = *directoryAsesorParent

	servicesAsesor, err := s.directoryTreeService.ListByParent(&directoryAsesor)

	if err != nil {
		return false, err
	}

	directoryInstitution.Id = *directoryInstitutionParentId

	servicesInstitution, err := s.directoryTreeService.ListByParent(&directoryInstitution)
	if err != nil {
		return false, err
	}

	institutionFoldersMap := make(map[string]*dtdto.GetDirectoryTreeRes)
	for _, folder := range servicesInstitution.Directory {
		institutionFoldersMap[folder.Name] = folder
	}

	var commonFolders [][]*dtdto.GetDirectoryTreeRes
	if servicesAsesor.Directory != nil {
		for _, folder := range servicesAsesor.Directory {
			if institutionFoldersMap[folder.Name] != nil {
				commonFolders = append(commonFolders, []*dtdto.GetDirectoryTreeRes{
					folder,
					institutionFoldersMap[folder.Name],
				})
			}
		}
	}

	for _, folder := range commonFolders {
		directoryAsesor.Id = folder[0].Id
		folderDocuments, err := s.directoryTreeService.ListByParent(&directoryAsesor)
		if err != nil {
			return false, err
		}
		if folderDocuments.Document == nil {
			continue
		}

		for _, doc := range *folderDocuments.Document {
			nameFolder := fmt.Sprintf("%s%s", asesorEnum, strconv.Itoa(int(directoryAsesor.AsesorID)))
			docBytes, err := digitalOcean.DownloadDocumentByte(doc.FileRute, nameFolder)
			if err != nil {
				return false, err
			}

			formatReq := ds.WordApplyReq{
				Docu:   docBytes,
				Header: convertToDocuValues(map[string]string{"Company Name": institution.InstitutionName}),
				Body:   convertToDocuValues(map[string]string{"Company Name": institution.InstitutionName}),
				Footer: convertToDocuValues(map[string]string{"Company Name": institution.InstitutionName}),
			}

			format, err := s.formatService.ApplyWordFormat(formatReq)
			if err != nil {
				return false, err
			}
			documentReq := ddto.CreateDocumentByteReq{
				DirectoryId:   folder[1].Id,
				Status:        "en revision",
				File:          &format,
				Filename:      doc.Name,
				InstitutionId: folder[1].InstitutionID,
			}

			_, err = s.documentsService.Create(&documentReq)
			if err != nil {
				return false, err
			}
		}
	}
	return true, nil
}

func convertToDocuValues(data map[string]string) []docustream.DocuValue {
	values := make([]docustream.DocuValue, 0, len(data))
	for k, v := range data {
		values = append(values, docustream.DocuValue{
			Key:   k,
			Value: v,
		})
	}
	return values
}
