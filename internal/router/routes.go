package router

import (
	"optitech/internal/handler"
)

func (s *Server) Routes() {
	r := s.app
	r.Group("/api")
	r.Post("/client", handler.CreateClientHandler)
	r.Get("/client/:id", handler.GetClientHandler)
}
