package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	document "optitech/internal/service/documents"
)

func (s *Server) RoutesDocument() {
	r := s.app
	repoDocument := repository.NewRepositoryDocument(&repository.Queries)
	service := document.NewServiceDocument(repoDocument)
	handler := handler.NewHandlerDocument(service)
	serviceRoute := r.Group("/api/document")
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/", handler.CreateDocument)
	serviceRoute.Delete("/:id", handler.DeleteDocument)
	serviceRoute.Get("download/:id", handler.DownloadDocumentById)
	serviceRoute.Put("name", handler.UpdateDocument)
}
