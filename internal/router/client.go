package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/client"
)

func (s *Server) RoutesClient() {
	r := s.app
	repoService := repository.NewRepositoryClient(&repository.Queries)
	sevice := service.NewServiceClient(repoService)
	handler := handler.NewHandlerClient(sevice)
	serviceRoute := r.Group("/api/client")
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/", handler.List)
	serviceRoute.Post("", handler.Create)

}
