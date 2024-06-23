package service

import (
	"optitech/database"
	dto "optitech/internal/dto/institution_client"
	"optitech/internal/repository"
	"testing"
	"time"

	sq "optitech/internal/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestClientClient(t *testing.T) {
	db, err := database.Connect()

	assert.NotNil(t, db)
	assert.Nil(t, err)

	repo := repository.NewRepositoryInstitutionClient(&repository.Queries)
	service := NewServiceInstitutionClient(repo)
	const InstitutionID = 3

	t.Run("Create institution service", func(t *testing.T) {
		req := &sq.CreateInstitutionClientParams{
			InstitutionID: InstitutionID,
			CreatedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
			ClientID:      3,
		}
		err = service.Create(&[]sq.CreateInstitutionClientParams{*req})
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

	t.Run("Get the all clients of institution", func(t *testing.T) {
		res, err := service.List(InstitutionID)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("Delete client by institution and service", func(t *testing.T) {
		err := service.DeleteById(&dto.GetInstitutionClientReq{ClientId: 3, InstitutionId: InstitutionID})

		assert.NotNil(t, err)
		assert.Nil(t, err)
	})
	t.Run("Exist client of institution", func(t *testing.T) {
		res := service.Exists(&sq.ExistsInstitutionClientParams{ClientID: 3, InstitutionID: InstitutionID})
		assert.Equal(t, res, true)
		assert.Nil(t, err)
	})

	t.Run("Recover service of institution", func(t *testing.T) {
		err := service.Recover(&sq.RecoverInstitutionClientParams{ClientID: 3, InstitutionID: InstitutionID})
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

	t.Run("Delete clients of institution", func(t *testing.T) {
		err := service.DeleteByInstitution(InstitutionID)
		assert.NotNil(t, err)
		assert.Nil(t, err)
	})

}
