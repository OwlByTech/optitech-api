package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/services"
)

func (s *Server) RoutesServices() {
	r := s.app
	repo_service := repository.NewRepositoryService(&repository.Queries)
	sevice := service.NewServiceServices(repo_service)
	handler := handler.NewHandlerService(sevice)
	service_route := r.Group("/api/service")
	service_route.Post("/", handler.Create)
	service_route.Get("/:id", handler.Get)
	service_route.Get("/", handler.List)

}
