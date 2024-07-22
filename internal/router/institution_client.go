package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/institution_client"
)

func (s *Server) RoutesInstitutionClient() {
	r := s.app
	repo := repository.NewRepositoryInstitutionClient(&repository.Queries)
	sevice := service.NewServiceInstitutionClient(repo)
	handler := handler.NewHandlerInstitutionClient(sevice)
	serviceRoute := r.Group("/api/institution-client")
	serviceRoute.Post("/", handler.Update)

}
