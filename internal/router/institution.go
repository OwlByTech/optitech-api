package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/institution"
	institutionClient "optitech/internal/service/institution_client"
	serviceInstitution "optitech/internal/service/institution_services"
)

func (s *Server) RoutesInstitution() {
	r := s.app
	repositoryInstitutionService := repository.NewRepositoryInstitutionServices(&repository.Queries)
	serviceInstitutionService := serviceInstitution.NewServiceInstitutionServices(repositoryInstitutionService)
	repositoryInstitutionClient := repository.NewRepositoryInstitutionClient(&repository.Queries)
	serviceInstitutionClient := institutionClient.NewServiceInstitutionClient(repositoryInstitutionClient)
	repositoryInstitution := repository.NewRepositoryInstitution(&repository.Queries)
	serviceInstitution := service.NewServiceInstitution(repositoryInstitution, serviceInstitutionService, serviceInstitutionClient)
	handler := handler.NewHandlerInstitution(serviceInstitution)
	institutionRoute := r.Group("/api/institution")
	institutionRoute.Post("/", handler.Create)
	institutionRoute.Get("/:id", handler.Get)
	institutionRoute.Get("/", handler.List)

}
