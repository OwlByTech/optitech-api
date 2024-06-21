package seeders

import (
	"context"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	sdto "optitech/internal/dto/standards"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StandardUp(fileName string) error {
	ctx := context.Background()
	var curTime time.Time

	var standards []sdto.CreateStandardsReq
	err := json_reader.ReadFromJSON(fileName, &standards)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	curTime = time.Now()

	var sqStandard []sq.CreateStandardParams
	for _, data := range standards {
		standard := sq.CreateStandardParams{
			ServiceID:  data.ServiceId,
			Name:       data.Name,
			Complexity: pgtype.Text{String: data.Complexity},
			Modality:   data.Modality,
			Article:    data.Article,
			Section:    data.Section,
			Paragraph:  pgtype.Text{String: data.Paragraph},
			Criteria:   data.Criteria,
			Comply:     pgtype.Bool{Bool: data.Comply},
			Applys:     pgtype.Bool{Bool: data.Applys},
			CreatedAt: pgtype.Timestamp{
				Time:  curTime,
				Valid: true,
			},
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
	deleteAt := pgtype.Timestamp{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllStandards(ctx, deleteAt)
	if err != nil {
		return err
	}

	_ = r.RowsAffected()

	log.Printf("Standard Down seeder run successfully")
	return nil
}
