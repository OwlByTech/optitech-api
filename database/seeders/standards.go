package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	sdto "optitech/internal/dto/standards"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"
)

func StandardUp(fileName string) error {
	ctx := context.Background()
	curTime := time.Now()

	var standards []sdto.CreateStandardsReq
	err := json_reader.ReadFromJSON(fileName, &standards)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	var sqStandard []sq.CreateStandardParams
	for _, data := range standards {
		standard := sq.CreateStandardParams{
			ServiceID:  data.ServiceId,
			Standard:   data.Standard,
			Complexity: sql.NullString{String: data.Complexity},
			Modality:   data.Modality,
			Article:    data.Article,
			Section:    data.Section,
			Paragraph:  sql.NullString{String: data.Paragraph},
			Criteria:   data.Criteria,
			Comply:     sql.NullBool{Bool: data.Comply},
			Applys:     sql.NullBool{Bool: data.Applys},
			CreatedAt:  curTime,
		}
		sqStandard = append(sqStandard, standard)
	}

	for _, standard := range sqStandard {
		if _, err := repository.Queries.CreateStandard(ctx, standard); err != nil {
			return fmt.Errorf("error inserting data in db: %v", err)
		}
	}

	log.Printf("Standard Up seeder run successfully")
	return nil
}

func StandardDown() error {
	ctx := context.Background()
	curTime := time.Now()
	deleteAt := sql.NullTime{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllStandards(ctx, deleteAt)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Standard Down seeder run successfully")
	return nil
}
