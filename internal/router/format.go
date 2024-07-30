package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	format "optitech/internal/service/format"
)

func (s *Server) RoutesFormat() {
	r := s.app
	repoFormat := repository.NewRepositoryFormat(&repository.Queries)
	service := format.NewServiceFormat(repoFormat)
	handler := handler.NewHandlerFormat(service)
	serviceRoute := r.Group("/api/format")
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/create", handler.Create)
}
