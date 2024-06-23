package seeders

import (
	"context"
	"fmt"
	"log"
	json_reader "optitech/database/json_data"
	rdto "optitech/internal/dto/role_permission"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func RolePermissionUp(fileName string) error {
	ctx := context.Background()
	var curTime time.Time

	var rolePermissions []rdto.CreateRolePermissionReq
	err := json_reader.ReadFromJSON(fileName, &rolePermissions)
	if err != nil {
		return fmt.Errorf("error reading json %v", err)
	}

	curTime = time.Now()

	var sqRolePermission []sq.CreateRolePermissionParams
	for _, data := range rolePermissions {
		rolePermission := sq.CreateRolePermissionParams{
			PermissionID: data.PermissionId,
			RoleID:       data.RoleId,
			CreatedAt: pgtype.Timestamp{
				Time:  curTime,
				Valid: true,
			},
		}
		sqRolePermission = append(sqRolePermission, rolePermission)
	}

	for _, rolePermission := range sqRolePermission {
		if _, err := repository.Queries.CreateRolePermission(ctx, rolePermission); err != nil {
			return fmt.Errorf("error inserting data in db: %v", err)
		}
	}

	return nil
}

func RolePermissionDown() error {
	ctx := context.Background()
	curTime := time.Now()
	deleteAt := pgtype.Timestamp{
		Time:  curTime,
		Valid: true,
	}
	r, err := repository.Queries.DeleteAllRolePermissions(ctx, deleteAt)
	if err != nil {
		return err
	}

	_ = r.RowsAffected()

	log.Printf("Role Permission Down seeder run successfully")
	return nil
}
