package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	directoryTreeService "optitech/internal/service/directory_tree"
	documentsService "optitech/internal/service/documents"
	service "optitech/internal/service/institution"
	institutionClient "optitech/internal/service/institution_client"
	serviceInstitution "optitech/internal/service/institution_services"
	services "optitech/internal/service/services"
)

var repositoryInstitutionService = repository.NewRepositoryInstitutionServices(&repository.Queries)
var serviceInstitutionService = serviceInstitution.NewServiceInstitutionServices(repositoryInstitutionService)
var repositoryInstitutionClient = repository.NewRepositoryInstitutionClient(&repository.Queries)
var serviceInstitutionClient = institutionClient.NewServiceInstitutionClient(repositoryInstitutionClient)
var repositoryInstitution = repository.NewRepositoryInstitution(&repository.Queries)
var repositoryDirectoryTree = repository.NewRepositoryDirectoryTree(&repository.Queries)
var repositoryDocuments = repository.NewRepositoryDocument(&repository.Queries)
var serviceDocument = documentsService.NewServiceDocument(repositoryDocuments)
var serviceDirectoryTree = directoryTreeService.NewServiceDirectory(repositoryDirectoryTree, serviceDocument)
var repositoryServices = repository.NewRepositoryService(&repository.Queries)
var serviceServices = services.NewServiceServices(repositoryServices)
var ServiceInstitution = service.NewServiceInstitution(repositoryInstitution, serviceInstitutionService, serviceInstitutionClient, serviceDirectoryTree, serviceServices)

func (s *Server) RoutesInstitution() {
	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}
	r := s.app

	handler := handler.NewHandlerInstitution(ServiceInstitution)
	institutionRoute := r.Group("/api/institution")
	institutionRoute.Post("/", clientMiddleware.ClientJWT, handler.Create)
	institutionRoute.Get("/:id", handler.Get)
	institutionRoute.Get("/", clientMiddleware.ClientJWT, handler.GetByClient)
	institutionRoute.Get("/all", handler.List)
	institutionRoute.Delete("/:id", handler.Delete)
	institutionRoute.Put("/:id", handler.Update)
	institutionRoute.Post("/logo/:id", clientMiddleware.ClientJWT, handler.UpdateLogo)
	institutionRoute.Get("/logo/:id", clientMiddleware.ClientJWT, handler.GetLogo)
	institutionRoute.Post("/asesor", handler.UpdateAsesor)

}
