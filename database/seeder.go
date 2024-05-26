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
	if err := seeders.ClientUp(); err != nil {
		return err
	}
	if err := seeders.AsesorUp(); err != nil {
		return err
	}
	if err := seeders.InstitutionUp(); err != nil {
		return err
	}
	if err := seeders.FormatUp(); err != nil {
		return err
	}
	if err := seeders.DocumentsUp(); err != nil {
		return err
	}
	if err := seeders.DocumentClientUp(); err != nil {
		return err
	}
	if err := seeders.InstitutionClientUp(); err != nil {
		return err
	}
	return nil
}

type SeederDown struct{}

func (s SeederDown) Execute() error {
	if err := seeders.FormatDown(); err != nil {
		return err
	}
	if err := seeders.AsesorDown(); err != nil {
		return err
	}
	if err := seeders.ClientDown(); err != nil {
		return err
	}
	if err := seeders.InstitutionDown(); err != nil {
		return err
	}
	if err := seeders.DocumentsDown(); err != nil {
		return err
	}
	if err := seeders.DocumentClientDown(); err != nil {
		return err
	}
	if err := seeders.InstitutionClientDown(); err != nil {
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
