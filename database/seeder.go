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

func (s SeederUp) Execute() error {
	return nil
}

type SeederDown struct{}

func (s SeederDown) Execute() error {
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
