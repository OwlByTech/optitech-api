package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	dts "optitech/internal/service/directory_tree"
	document "optitech/internal/service/documents"
)

func (s *Server) RoutesDirectoryTree() {
	r := s.app
	repoDocument := repository.NewRepositoryDocument(&repository.Queries)
	serviceDocument := document.NewServiceDocument(repoDocument)
	repoService := repository.NewRepositoryDirectoryTree(&repository.Queries)
	service := dts.NewServiceDirectory(repoService, serviceDocument)
	handler := handler.NewHandlerDirectoryTree(service)

	serviceRoute := r.Group("/api/directory-tree")
	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/parent/:id", handler.ListByParent)
	serviceRoute.Get("/child/:id", handler.ListByChild)
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/route/:id", handler.GetRoute)
	serviceRoute.Post("/", handler.Create)
	serviceRoute.Delete("/delete/:id", handler.Delete)
}
