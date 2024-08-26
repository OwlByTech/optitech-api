package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	dts "optitech/internal/service/directory_tree"
)

var repoService = repository.NewRepositoryDirectoryTree(&repository.Queries)
var serviceDirectoryTree = dts.NewServiceDirectory(repoService, serviceDocument)

func (s *Server) RoutesDirectoryTree() {
	r := s.app
	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}
	directoryMiddleware := middleware.DirectoryMiddleware{
		InstitutionService: ServiceInstitution,
		AsesorService:      serviceAsesor,
	}
	handler := handler.NewHandlerDirectoryTree(serviceDirectoryTree)

	serviceRoute := r.Group("/api/directory-tree")
	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/parent/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.ListByParent)
	serviceRoute.Get("/child/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.ListByChild)
	serviceRoute.Get("/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.Get)
	serviceRoute.Get("/route/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.GetRoute)
	serviceRoute.Post("/", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.Create)
	serviceRoute.Delete("/delete/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.Delete)
	serviceRoute.Put("/update/:id", clientMiddleware.ClientJWT, directoryMiddleware.DirectoryJWT, handler.Update)
}
