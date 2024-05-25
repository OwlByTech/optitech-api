package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

func FormatUp() error {

	ctx := context.Background()
	curTime := time.Now()
	format := sq.CreateFormatParams{
		AsesorID:    1,
		FormatName:  "Formato de manuales",
		Description: "formato según documento de elaboración",
		Items:       []string{"cabecera", "parrafo", "footer"},
		Extension:   sq.ExtensionsPdf,
		Version:     "Version 1",
		CreateAt:    curTime,
	}

	_, err := repository.Queries.CreateFormat(ctx, format)
	if err != nil {
		return err
	}

	log.Printf("Format Up seeder run successfully")
	return nil
}

func FormatDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllFormats(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Format Down seeder run successfully")
	return nil
}
