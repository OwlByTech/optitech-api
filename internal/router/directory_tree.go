package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	dts "optitech/internal/service/directory_tree"
)

func (s *Server) RoutesDirectoryClient() {
	r := s.app
	repoService := repository.NewRepositoryDirectoryTree(&repository.Queries)
	service := dts.NewServicDirectory(repoService)
	handler := handler.NewHnadlerDirectoryTree(service)
	serviceRoute := r.Group("/api/directory-tree")

	serviceRoute.Get("/", handler.Get)
	serviceRoute.Post("/", handler.Create)

}
