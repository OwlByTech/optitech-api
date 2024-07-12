package router

import (
	"optitech/internal/handler"
	"optitech/internal/middleware"
)

func (s *Server) RoutesClient() {
	r := s.app

	handler := handler.NewHandlerClient(SeviceClient)
	serviceRoute := r.Group("/api/client")

	// We should initialize all the middlewares here
	clientMiddleware := middleware.ClientMiddleware{
		ClientService: SeviceClient,
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
	serviceRoute.Put("/update/:id", handler.Update)
	serviceRoute.Post("/photo/:id", handler.UpdatePhoto)
	serviceRoute.Post("/status", handler.UpdateStatus)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
	serviceRoute.Post("/reset-password", handler.ResetPassword)
	serviceRoute.Post("/reset-password-token", handler.ResetPasswordToken)
	serviceRoute.Get("/validate/reset-password-token", handler.ValidateResetPasswordToken)
}
