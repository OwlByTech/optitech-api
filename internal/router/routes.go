package router

import (
	"optitech/internal/handler"
)

func (s Server) AttachRoutes() {
	r := s.app

	r.Get("/", handler.GetClientHandler)
}
