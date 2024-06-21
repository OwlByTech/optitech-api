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

	// We should initialize all the middlewares here

	// TODO: protect the routes with middlewares
	serviceRoute.Get("/:id", handler.Get)
	serviceRoute.Get("/", handler.List)
	serviceRoute.Post("", handler.Create)
	serviceRoute.Put("/update/:id", handler.Update)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
	serviceRoute.Post("/reset-password", handler.ResetPassword)
	serviceRoute.Post("/reset-password-token", handler.ResetPasswordToken)
}
