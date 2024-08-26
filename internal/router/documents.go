package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	document "optitech/internal/service/documents"
)

var repoDocument = repository.NewRepositoryDocument(&repository.Queries)
var serviceDocument = document.NewServiceDocument(repoDocument)

func (s *Server) RoutesDocument() {

	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}
	directoryMiddleware := middleware.DirectoryMiddleware{
		InstitutionService: ServiceInstitution,
		AsesorService:      serviceAsesor,
	}
	institutionMiddleware := middleware.InstitutionMiddleware{
		InstitutionService: ServiceInstitution,
	}

	r := s.app

	handler := handler.NewHandlerDocument(serviceDocument)
	serviceRoute := r.Group("/api/document")
	serviceRoute.Get("/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.Get)
	serviceRoute.Post("/", clientMiddleware.ClientJWT, institutionMiddleware.InstitutionJWT, handler.CreateDocument)
	serviceRoute.Delete("/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.DeleteDocument)
	serviceRoute.Get("download/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.DownloadDocumentById)
	serviceRoute.Put("update", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.UpdateDocument)
	serviceRoute.Put("status", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.UpdateStatusById)
	serviceRoute.Post("version", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.CreateVersion)
}
