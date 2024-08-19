package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	format "optitech/internal/service/format"
)

func (s *Server) RoutesFormat() {

	repoFormat := repository.NewRepositoryFormat(&repository.Queries)
	service := format.NewServiceFormat(repoFormat, serviceDocument)
	handler := handler.NewHandlerFormat(service)

	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
	}

	asesorMiddleware := middleware.AsesorMiddleware{
		AsesorService: serviceAsesor,
	}

	r := s.app
	serviceRoute := r.Group("/api/format")
	serviceRoute.Get("/all", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.List)
	serviceRoute.Get("/:id", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.Get)
	serviceRoute.Post("/", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.Create)
	serviceRoute.Post("/listId", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.ListById)
	serviceRoute.Delete("/:id", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.Delete)
	serviceRoute.Put("/:id", clientMiddleware.ClientJWT, asesorMiddleware.AsesorJWT, handler.Update)
}
