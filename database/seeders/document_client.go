package seeders

import (
	"context"
	"log"
	"optitech/internal/repository"
	"time"

	sq "optitech/internal/sqlc"
)

// TODO: Add a interface for all seeders
func DocumentClientUp() error {
	ctx := context.Background()
	curTime := time.Now()
	document_client := sq.CreateDocumentClientParams{
		ClientID:   2,
		DocumentID: 1,
		Action:     sq.ActionCreado,
		CreateAt:   curTime,
	}

	_, err := repository.Queries.CreateDocumentClient(ctx, document_client)
	if err != nil {
		return err
	}

	log.Printf("Document Client Up seeder run successfully")
	return nil
}

func DocumentClientDown() error {
	ctx := context.Background()
	r, err := repository.Queries.DeleteAllDocumentClients(ctx)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Document Client Down seeder run successfully")
	return nil
}
