package seeders

import (
	"context"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	sdto "optitech/internal/dto/services"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ServiceUp(fileName string) error {
	ctx := context.Background()
	var curTime time.Time

	var services []sdto.CreateServiceReq
	err := json_reader.ReadFromJSON(fileName, &services)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	curTime = time.Now()

	var sqServices []sq.CreateServicesParams
	for _, data := range services {
		service := sq.CreateServicesParams{
			Name:      data.Name,
			CreatedAt: pgtype.Timestamp{Time: curTime, Valid: true},
		}
		sqServices = append(sqServices, service)
	}

	for _, service := range sqServices {
		if _, err := repository.Queries.CreateServices(ctx, service); err != nil {
			return fmt.Errorf("error inserting data in db: %v", err)
		}
	}

	log.Printf("Service Up seeder run successfully")
	return nil
}

func ServiceDown() error {
	ctx := context.Background()
	curTime := time.Now()
	deleteAt := pgtype.Timestamp{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllServicess(ctx, deleteAt)
	if err != nil {
		return err
	}

	_ = r.RowsAffected()

	log.Printf("Service Down seeder run successfully")
	return nil
}
