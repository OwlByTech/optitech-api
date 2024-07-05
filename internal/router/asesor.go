package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/asesor"
)

func (s *Server) RoutesAsesor() {
	r := s.app
	repoService := repository.NewRepositoryAsesor(&repository.Queries)
	sevice := service.NewServiceAsesor(repoService)
	handler := handler.NewHandlerAsesor(sevice)
	serviceRoute := r.Group("/api/asesor")

	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/", handler.Create)
	serviceRoute.Put("/:id", handler.Update)
	serviceRoute.Delete("/:id", handler.Delete)
}
