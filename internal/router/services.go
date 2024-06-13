package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/services"
)

func (s *Server) RoutesServices() {
	r := s.app
	repoService := repository.NewRepositoryService(&repository.Queries)
	sevice := service.NewServiceServices(repoService)
	handler := handler.NewHandlerService(sevice)
	serviceRoute := r.Group("/api/services")
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/", handler.List)

}
