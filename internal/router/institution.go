package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	service "optitech/internal/service/institution"
	institutionClient "optitech/internal/service/institution_client"
	serviceInstitution "optitech/internal/service/institution_services"
)

func (s *Server) RoutesInstitution() {
	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}
	r := s.app
	repositoryInstitutionService := repository.NewRepositoryInstitutionServices(&repository.Queries)
	serviceInstitutionService := serviceInstitution.NewServiceInstitutionServices(repositoryInstitutionService)
	repositoryInstitutionClient := repository.NewRepositoryInstitutionClient(&repository.Queries)
	serviceInstitutionClient := institutionClient.NewServiceInstitutionClient(repositoryInstitutionClient)
	repositoryInstitution := repository.NewRepositoryInstitution(&repository.Queries)
	serviceInstitution := service.NewServiceInstitution(repositoryInstitution, serviceInstitutionService, serviceInstitutionClient)
	handler := handler.NewHandlerInstitution(serviceInstitution)
	institutionRoute := r.Group("/api/institution")
	institutionRoute.Post("/", clientMiddleware.ClientJWT, handler.Create)
	institutionRoute.Get("/:id", handler.Get)
	institutionRoute.Get("/", clientMiddleware.ClientJWT, handler.GetByClient)
	institutionRoute.Get("/all", handler.List)
	institutionRoute.Delete("/:id", handler.Delete)
	institutionRoute.Put("/:id", handler.Update)
	institutionRoute.Put("/logo/:id", clientMiddleware.ClientJWT, handler.UpdateLogo)
	institutionRoute.Post("/asesor", handler.UpdateAsesor)

}
