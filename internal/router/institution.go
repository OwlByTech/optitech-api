package router

import (
	"optitech/internal/handler"
)

func (s Server) AttachRoutes() {
	r := s.app
	institution_route := r.Group("/api/institution")
	institution_route.Post("/", handler.CreateInstitutionHandler)
	institution_route.Put("/:id", handler.UpdateInstitutionHandler)
	institution_route.Delete("/:id", handler.DeleteInstitutionHandler)
	institution_route.Get("/:id", handler.GetInstitutionHandler)
	institution_route.Get("/", handler.ListInstitutionsHandler)
	institution_route.Delete("/:id", handler.DeleteInstitutionHandler)

}
