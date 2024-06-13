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
	return seeders.PermissionUp("database/json_data/permission.json")
}

type SeederDown struct{}

func (s SeederDown) Execute() error {
	return seeders.PermissionDown()
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
