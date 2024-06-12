package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	rdto "optitech/internal/dto/roles"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"
)

func RoleUp(fileName string) error {
	ctx := context.Background()
	var curTime time.Time

	var role []rdto.CreateRoleReq
	err := json_reader.ReadFromJSON(fileName, &role)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	curTime = time.Now()

	var sqRoles []sq.CreateRoleParams
	for _, data := range role {
		role := sq.CreateRoleParams{
			RoleName:    data.RoleName,
			Description: data.Description,
			CreatedAt:   curTime,
		}
		sqRoles = append(sqRoles, role)
	}

	for _, role := range sqRoles {
		if _, err := repository.Queries.CreateRole(ctx, role); err != nil {
			return fmt.Errorf("error inserting data in db: %v", err)
		}
	}

	log.Printf("Role Up seeder run successfully")
	return nil
}

func RoleDown() error {
	ctx := context.Background()
	curTime := time.Now()
	deleteAt := sql.NullTime{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllRoles(ctx, deleteAt)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Role Down seeder run successfully")
	return nil
}
