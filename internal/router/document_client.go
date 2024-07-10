package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	documentClient "optitech/internal/service/document_client"
)

func (s *Server) RoutesDocumentClient() {
	r := s.app
	repoDocumentClient := repository.NewRepositoryDocumentClient(&repository.Queries)
	service := documentClient.NewServiceDocumentClient(repoDocumentClient)
	handler := handler.NewhandlerDocumentClient(service)
	serviceRoute := r.Group("/api/document-client")
	serviceRoute.Get("/:id", handler.Get)
}
