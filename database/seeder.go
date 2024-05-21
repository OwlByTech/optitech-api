package database

import (
	"fmt"
	"log"
	"optitech/database/seeders"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
)

func Seeder(arg string) error {
	db, err := Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}

	repository.Queries = *sq.New(db)

	switch arg {
	case "up":
		if err := seederUp(); err != nil {
			return err
		}

	case "down":
		if err := seederDown(); err != nil {
			return err
		}

	default:
		return fmt.Errorf("You must provide up or down arguments")
	}

	return nil
}

func seederUp() error {
	if err := seeders.ClientUp(); err != nil {
		return err
	}
	if err := seeders.AsesorUp(); err != nil {
		return err
	}
	if err := seeders.InstitutionUp(); err != nil {
		return err
	}
	return nil
}

func seederDown() error {
	if err := seeders.ClientDown(); err != nil {
		return err
	}
	if err := seeders.AsesorDown(); err != nil {
		return err
	}
	if err := seeders.InstitutionDown(); err != nil {
		return err
	}
	return nil
}
