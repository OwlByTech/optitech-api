package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders

func AsesorUp() error {
	ctx := context.Background()
	curTime := time.Now()
	asesor := sq.CreateAsesorParams{
		Username: "JosephSC0121",
		About:    "Hola mundo ",
		CreateAt: curTime,
	}

	_, err := repository.Queries.CreateAsesor(ctx, asesor)
	if err != nil {
		return err
	}

	log.Printf("Asesor Up seeder run successfully")
	return nil
}

func AsesorDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllAsesors(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Asesor Down seeder run successfully")
	return nil
}
