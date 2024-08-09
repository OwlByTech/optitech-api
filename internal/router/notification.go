package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	dts "optitech/internal/service/notification"
)

func (s *Server) RoutesNotification() {
	r := s.app

	repoService := repository.NewrepositoryNotification(&repository.Queries)
	service := dts.NewServiceNotification(repoService)
	hanlder := handler.NewHandlerNotification(service)

	serviceRoute := r.Group("/api/notification")

	serviceRoute.Post("/", hanlder.Create)
}
