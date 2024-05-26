package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders

func DocumentsUp() error {
	ctx := context.Background()
	curTime := time.Now()
	documents := sq.CreateDocumentParams{
		FormatID:      1,
		InstitutionID: 1,
		ClientID:      2,
		FileRute:      "/documets/file.doc",
		Status:        sq.StatusEnrevision,
		CreateAt:      curTime,
	}

	_, err := repository.Queries.CreateDocument(ctx, documents)
	if err != nil {
		return err
	}

	log.Printf("Documents Up seeder run successfully")
	return nil
}

func DocumentsDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllDocuments(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Documents Down seeder run successfully")
	return nil
}
