package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	service "optitech/internal/service/client"
)

func (s *Server) RoutesClient() {
	r := s.app
	repoService := repository.NewRepositoryClient(&repository.Queries)
	sevice := service.NewServiceClient(repoService)
	handler := handler.NewHandlerClient(sevice)
	serviceRoute := r.Group("/api/client")

	// We should initialize all the middlewares here
	clientMiddleware := &middleware.ClientMiddleware{
		ClientService: sevice,
	}

	// TODO: protect the routes with middlewares
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/", clientMiddleware.ClientJWT, handler.List)
	serviceRoute.Post("", handler.Create)
	serviceRoute.Put("/update/:id", handler.Update)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
}
