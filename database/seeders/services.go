package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	sdto "optitech/internal/dto/services"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"
)

func ServiceUp(fileName string) error {
	ctx := context.Background()
	curTime := time.Now()

	var services []sdto.CreateServiceReq
	err := json_reader.ReadFromJSON(fileName, &services)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	var sqServices []sq.CreateServicesParams
	for _, data := range services {
		service := sq.CreateServicesParams{
			ServiceName: data.Name,
			CreatedAt:   curTime,
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
	deleteAt := sql.NullTime{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllServicess(ctx, deleteAt)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Service Down seeder run successfully")
	return nil
}
