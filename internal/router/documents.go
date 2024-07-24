package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	document "optitech/internal/service/documents"
)

func (s *Server) RoutesDocument() {

	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}
	r := s.app

	repoDocument := repository.NewRepositoryDocument(&repository.Queries)
	service := document.NewServiceDocument(repoDocument)
	handler := handler.NewHandlerDocument(service)
	serviceRoute := r.Group("/api/document")
	serviceRoute.Get("/:id", clientMiddleware.ClientJWT, handler.Get)
	serviceRoute.Post("/", clientMiddleware.ClientJWT, handler.CreateDocument)
	serviceRoute.Delete("/:id", clientMiddleware.ClientJWT, handler.DeleteDocument)
	serviceRoute.Get("download/:id", clientMiddleware.ClientJWT, handler.DownloadDocumentById)
	serviceRoute.Put("update", clientMiddleware.ClientJWT, handler.UpdateDocument)
}
