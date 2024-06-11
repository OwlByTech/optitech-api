package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/institution"
	institution_client "optitech/internal/service/institution_client"
	service_institution "optitech/internal/service/institution_services"
)

func (s *Server) RoutesInstitution() {
	r := s.app
	repo_institution_services := repository.NewRepositoryInstitutionServices(&repository.Queries)
	service_institution_services := service_institution.NewServiceInstitutionServices(repo_institution_services)
	repo_institution_client := repository.NewRepositoryInstitutionClient(&repository.Queries)
	service_institution_client := institution_client.NewServiceInstitutionClient(repo_institution_client)
	repo_institution := repository.NewRepositoryInstitution(&repository.Queries)
	sevice_institution := service.NewServiceInstitution(repo_institution, service_institution_services, service_institution_client)
	handler := handler.NewHandlerInstitution(sevice_institution)
	institution_route := r.Group("/api/institution")
	institution_route.Post("/", handler.Create)
	institution_route.Get("/:id", handler.Get)
	institution_route.Get("/", handler.List)

}
