package database

import (
	"fmt"
	"log"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
)

type SeederStrategy interface {
	Execute() error
}

type SeederUp struct{}

type SeederDown struct{}

func Seeder(arg string) error {
	db, err := Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository.Queries = *sq.New(db)

	var strategy SeederStrategy

	switch arg {
	case "up":
	case "down":
	default:
		return fmt.Errorf("You must provide up or down arguments")
	}

	return strategy.Execute()
}
