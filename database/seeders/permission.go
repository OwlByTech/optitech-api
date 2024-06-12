package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	pdto "optitech/internal/dto/permission"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"
)

func PermissionUp(fileName string) error {
	ctx := context.Background()
	var curTime time.Time

	var permissions []pdto.CreatePermissionReq
	err := json_reader.ReadFromJSON(fileName, &permissions)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	curTime = time.Now()

	var sqPermissions []sq.CreatePermissionParams
	for _, data := range permissions {
		permission := sq.CreatePermissionParams{
			Name:        data.Name,
			Code:        data.Code,
			Description: data.Description,
			CreatedAt:   curTime,
		}
		sqPermissions = append(sqPermissions, permission)
	}

	for _, permission := range sqPermissions {
		if _, err := repository.Queries.CreatePermission(ctx, permission); err != nil {
			return fmt.Errorf("error inserting data in db: %v", err)
		}
	}

	log.Printf("Permission Up seeder run successfully")
	return nil
}

func PermissionDown() error {
	ctx := context.Background()
	curTime := time.Now()
	deleteAt := sql.NullTime{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllPermissions(ctx, deleteAt)
	if err != nil {
		return err
	}

	_, err = r.RowsAffected()

	if err != nil {
		return err
	}

	log.Printf("Pemission Down seeder run successfully")
	return nil
}
