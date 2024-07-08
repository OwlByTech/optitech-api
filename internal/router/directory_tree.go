package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	dts "optitech/internal/service/directory_tree"
)

func (s *Server) RoutesDirectoryTree() {
	r := s.app
	repoService := repository.NewRepositoryDirectoryTree(&repository.Queries)
	service := dts.NewServiceDirectory(repoService)
	handler := handler.NewHandlerDirectoryTree(service)
	serviceRoute := r.Group("/api/directory-tree")

	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/", handler.Create)
	serviceRoute.Delete("/delete/:id", handler.Delete)
}
