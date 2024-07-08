package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	drs "optitech/internal/service/directory_role"
)

func (s *Server) RoutesDirectoryRole() {
	r := s.app
	repoService := repository.NewRepositoryDirectoryRole(&repository.Queries)
	service := drs.NewServiceDirectoryRole(repoService)
	handler := handler.NewHandlerDirectoryRole(service)
	serviceRoute := r.Group("/api/directory-role")

	serviceRoute.Post("/", handler.Create)
}
