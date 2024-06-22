package router

import (
	"log"
	"optitech/internal/handler"
	"optitech/internal/middleware"
	"optitech/internal/repository"
	"optitech/internal/service"
	cs "optitech/internal/service/client"
	"optitech/internal/sqlc"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Server) RoutesClient() {
	r := s.app
	repoService := repository.NewRepositoryClient(&repository.Queries)
	sevice := cs.NewServiceClient(repoService)
	handler := handler.NewHandlerClient(sevice)
	serviceRoute := r.Group("/api/client")

	repoRolePermission := repository.NewRepositoryRolePermission(&repository.Queries)
	serviceRolePermission := service.NewServiceRolePermission(repoRolePermission)

	repoClientRole := repository.NewRepositoryClientRole(&repository.Queries)
	serviceClientRole := service.NewServiceClientRole(repoClientRole)

	serviceRoute.Get("/test", func(c *fiber.Ctx) error {
		now := time.Now()
		serviceClientRole.Create(&sqlc.CreateClientRoleParams{
			ClientID:  1,
			RoleID:    5,
			CreatedAt: pgtype.Timestamp{Time: now, Valid: true},
		})

		res, err := serviceClientRole.ListByClientId(1)
		if err != nil {
			log.Println(err)
		}

		log.Println(res)

		res2, err := serviceRolePermission.GetByRoleId(int32((*res)[6].Id))

		if err != nil {
			log.Println(err)
		}

		c.JSON(res2)

		return nil
	})

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
	serviceRoute.Put("/update/:id", handler.Update)
	serviceRoute.Delete("/delete/:id", handler.Delete)
	serviceRoute.Post("/login", handler.Login)
	serviceRoute.Post("/reset-password", handler.ResetPassword)
	serviceRoute.Post("/reset-password-token", handler.ResetPasswordToken)
	serviceRoute.Get("/validate/reset-password-token", handler.ValidateResetPasswordToken)
}
