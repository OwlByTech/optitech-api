package router

import (
	"fmt"
	"optitech/internal/repository"
	cs "optitech/internal/service/client"
	service "optitech/internal/service/client_role"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	Port uint16
}

func (s *Server) New() {
	s.app = fiber.New()
}

var repoServiceClientRole = repository.NewRepositoryClientRole(&repository.Queries)
var serviceClientRole = service.NewServiceClientRole(repoServiceClientRole)
var repoServiceClient = repository.NewRepositoryClient(&repository.Queries)
var SeviceClient = cs.NewServiceClient(repoServiceClient, serviceClientRole)

func (s *Server) ListenAndServe() error {
	s.RoutesClient()
	s.RoutesAsesor()
	s.RoutesServices()
	s.RoutesInstitution()
	s.RoutesInstitutionClient()
	s.RoutesFile()
	s.RoutesFormat()
	s.RoutesDocument()
	s.RoutesDirectoryTree()
	s.RoutesDocumentClient()
	s.RoutesDirectoryRole()
	s.RoutesNotification()
	err := s.app.Listen(fmt.Sprintf(":%d", s.Port))

	if err != nil {
		return err
	}

	return nil
}
