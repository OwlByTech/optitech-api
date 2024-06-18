package database

import (
	"fmt"
	"log"
	"optitech/database/seeders"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
)

type SeederStrategy interface {
	Execute() error
}

type SeederUp struct{}

func (s SeederUp) Execute() error {
	if err := seeders.PermissionUp("database/json_data/permission.json"); err != nil {
		return err
	}
	if err := seeders.RoleUp("database/json_data/role.json"); err != nil {
		return err
	}
	if err := seeders.ServiceUp("database/json_data/services.json"); err != nil {
		return err
	}
	if err := seeders.StandardUp("database/json_data/standards.json"); err != nil {
		return err
	}
	if err := seeders.RolePermissionUp("database/json_data/role_permission.json"); err != nil {
		return err
	}
	return nil
}

type SeederDown struct{}

func (s SeederDown) Execute() error {
	if err := seeders.PermissionDown(); err != nil {
		return err
	}
	if err := seeders.RoleDown(); err != nil {
		return err
	}
	if err := seeders.ServiceDown(); err != nil {
		return err
	}
	if err := seeders.StandardDown(); err != nil {
		return err
	}
	if err := seeders.RolePermissionDown(); err != nil {
		return err
	}
	return nil
}
func Seeder(arg string) error {
	db, err := Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository.Queries = *sq.New(db)

	var strategy SeederStrategy

	switch arg {
	case "up":
		strategy = SeederUp{}
	case "down":
		strategy = SeederDown{}
	default:
		return fmt.Errorf("You must provide up or down arguments")
	}

	return strategy.Execute()
}
