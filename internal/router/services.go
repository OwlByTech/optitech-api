package router

import (
	"optitech/internal/handler"
)

func (s *Server) RoutesServices() {
	r := s.app
	institution_route := r.Group("/api/services")
	institution_route.Post("/", handler.CreateServiceHandler)
	institution_route.Get("/:id", handler.GetServiceHandler)
	institution_route.Get("/", handler.ListServicesHandler)

}
