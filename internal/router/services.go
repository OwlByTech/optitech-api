package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/services"
)

var repoServiceServices = repository.NewRepositoryService(&repository.Queries)
var serviceServices = service.NewServiceServices(repoServiceServices)

func (s *Server) RoutesServices() {
	r := s.app

	handler := handler.NewHandlerService(serviceServices)
	serviceRoute := r.Group("/api/services")
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/", handler.List)

}
