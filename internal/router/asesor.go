package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/asesor"
)

var repoAsesor = repository.NewRepositoryAsesor(&repository.Queries)
var serviceAsesor = service.NewServiceAsesor(repoAsesor, serviceDirectoryTree, serviceServices)

func (s *Server) RoutesAsesor() {
	r := s.app
	handler := handler.NewHandlerAsesor(serviceAsesor)
	serviceRoute := r.Group("/api/asesor")

	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/", handler.Create)
	serviceRoute.Put("/:id", handler.Update)
	serviceRoute.Delete("/:id", handler.Delete)
}
