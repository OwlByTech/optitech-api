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

	// The following routes must be use "clientId" locals to get
	// the current user
	// https://docs.gofiber.io/api/ctx/#locals
	serviceRoute.Get("/", clientMiddleware.ClientJWT, handler.GetSecure)

	// TODO: protect the following routes with middlewares
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/all", handler.List)
	serviceRoute.Post("", handler.Create)
	serviceRoute.Put("/update/:id", handler.Update)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
}
