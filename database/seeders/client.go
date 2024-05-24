package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders

func ClientUp() error {
	ctx := context.Background()
	curTime := time.Now()
	client_asesor := sq.CreateClientParams{
		Email:     "asesor@owlbytech.com",
		GivenName: "asesor",
		Password:  "password",
		Surname:   "asesor",
		CreatedAt: curTime,
	}
	client_institution := sq.CreateClientParams{
		Email:     "institution@owlbytech.com",
		GivenName: "institution",
		Password:  "password",
		Surname:   "ips",
		CreatedAt: curTime,
	}

	_, err := repository.Queries.CreateClient(ctx, client_asesor)
	if err != nil {
		return err
	}
	_, err = repository.Queries.CreateClient(ctx, client_institution)
	if err != nil {
		return err
	}

	log.Printf("Client Up seeder run successfully")
	return nil
}

func ClientDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllClients(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Client Down seeder run successfully")
	return nil
}
