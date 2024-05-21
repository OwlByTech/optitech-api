package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders

func InstitutionUp() error {
	ctx := context.Background()
	curTime := time.Now()
	institution := sq.CreateInstitutionParams{
		InstitutionName: "nombre institucion",
		Descrip:         "Soy una institucion",
		Services:        []string{"medicina", "otro"},
		CreateAt:        curTime,
	}

	_, err := repository.Queries.CreateInstitution(ctx, institution)
	if err != nil {
		return err
	}

	log.Printf("Institution Up seeder run successfully")
	return nil
}

func InstitutionDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllInstitutions(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Institution Down seeder run successfully")
	return nil
}
