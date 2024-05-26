package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders

func InstitutionClientUp() error {
	ctx := context.Background()
	curTime := time.Now()
	institution_client := sq.CreateInstitutionClientParams{
		ClientID:      2,
		InstitutionID: 1,
		VinculatedAt:  curTime,
	}

	_, err := repository.Queries.CreateInstitutionClient(ctx, institution_client)
	if err != nil {
		return err
	}

	log.Printf("Institution ClientDown Up seeder run successfully")
	return nil
}

func InstitutionClientDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllInstitutionClients(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Institution ClientDown Down seeder run successfully")
	return nil
}
