package service

import (
	"optitech/database"
	dto "optitech/internal/dto/institution_services"
	"optitech/internal/repository"
	"testing"
	"time"

	sq "optitech/internal/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestClientServices(t *testing.T) {
	db, err := database.Connect()

	assert.NotNil(t, db)
	assert.Nil(t, err)

	repo := repository.NewRepositoryInstitutionServices(&repository.Queries)
	service := NewServiceInstitutionServices(repo)
	const InstitutionID = 2

	t.Run("Create institution service", func(t *testing.T) {
		req := &sq.CreateInstitutionServicesParams{
			InstitutionID: InstitutionID,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
			ServiceID:     3,
		}
		err = service.Create(&[]sq.CreateInstitutionServicesParams{*req})
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

	t.Run("Get the all services of institution", func(t *testing.T) {
		res, err := service.List(InstitutionID)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("Delete services by institution and service", func(t *testing.T) {
		err := service.DeleteById(&dto.GetInstitutionServicesReq{ServiceID: 3, InstitutionID: InstitutionID})
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})
	t.Run("Exist services of institution", func(t *testing.T) {
		res := service.Exists(&sq.ExistsInstitutionServiceParams{ServiceID: 3, InstitutionID: InstitutionID})
		assert.Equal(t, res, true)
		assert.Nil(t, err)
	})

	t.Run("Recover service of institution", func(t *testing.T) {
		err := service.Recover(&sq.RecoverInstitutionServiceParams{ServiceID: 3, InstitutionID: InstitutionID})
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

	t.Run("Delete services of institution", func(t *testing.T) {
		err := service.DeleteByInstitution(InstitutionID)
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

}
