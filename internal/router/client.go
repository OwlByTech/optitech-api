package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	cs "optitech/internal/service/client"
)

func (s *Server) RoutesClient() {
	r := s.app
	repoService := repository.NewRepositoryClient(&repository.Queries)
	sevice := cs.NewServiceClient(repoService)
	handler := handler.NewHandlerClient(sevice)
	serviceRoute := r.Group("/api/client")

	// We should initialize all the middlewares here
	clientMiddleware := middleware.ClientMiddleware{
		ClientService: sevice,
	}

	// The following routes must be use "clientId" locals to get
	// the current user
	// https://docs.gofiber.io/api/ctx/#locals
	serviceRoute.Get("/", clientMiddleware.ClientJWT, handler.GetSecure)

	// TODO: protect the following routes with middlewares
	// The following route must be first than the /:id
	serviceRoute.Get("/all", handler.List)
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Post("/", handler.Create)
	serviceRoute.Put("/update/:id", clientMiddleware.ClientJWT, handler.Update)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
	serviceRoute.Post("/reset-password", handler.ResetPassword)
	serviceRoute.Post("/reset-password-token", handler.ResetPasswordToken)
	serviceRoute.Get("/validate/reset-password-token", handler.ValidateResetPasswordToken)
}
