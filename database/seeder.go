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
		if err := SeederUp(); err != nil {
			return err
		}

	case "down":
		if err := SeederDown(); err != nil {
			return err
		}

	default:
		return fmt.Errorf("You must provide up or down arguments")
	}

	return nil
}

func SeederUp() error {
	return seeders.ClientUp()
}

func SeederDown() error {
	return seeders.ClientDown()
}
