package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/institution"
)

func (s *Server) RoutesInstitution() {
	r := s.app
	repo := repository.NewRepositoryInstitution(&repository.Queries)
	sevice := service.NewServiceInstitution(repo)
	handler := handler.NewServiceInstitution(sevice)
	institution_route := r.Group("/api/institution")
	institution_route.Post("/", handler.Create)
	institution_route.Get("/:id", handler.Get)
	institution_route.Get("/", handler.List)

}
