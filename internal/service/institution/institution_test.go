package service

import (
	"optitech/database"
	dto "optitech/internal/dto/institution"
	"optitech/internal/repository"
	directoryTree "optitech/internal/service/directory_tree"
	serviceDocuments "optitech/internal/service/documents"
	institutionClient "optitech/internal/service/institution_client"
	serviceInstitutionTest "optitech/internal/service/institution_services"
	serviceService "optitech/internal/service/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	db, err := database.Connect()

	assert.NotNil(t, db)
	assert.Nil(t, err)
	repositoryInstitutionService := repository.NewRepositoryInstitutionServices(&repository.Queries)
	serviceInstitutionService := serviceInstitutionTest.NewServiceInstitutionServices(repositoryInstitutionService)
	repositoryInstitutionClient := repository.NewRepositoryInstitutionClient(&repository.Queries)
	serviceInstitutionClient := institutionClient.NewServiceInstitutionClient(repositoryInstitutionClient)
	repositoryDirectoryTree := repository.NewRepositoryDirectoryTree(&repository.Queries)
	repositoryDocuments := repository.NewRepositoryDocument(&repository.Queries)
	documentService := serviceDocuments.NewServiceDocument(repositoryDocuments)
	serviceDirectoryTree := directoryTree.NewServiceDirectory(repositoryDirectoryTree, documentService)
	repo := repository.NewRepositoryInstitution(&repository.Queries)
	repositoryServices := repository.NewRepositoryService(&repository.Queries)
	services := serviceService.NewServiceServices(repositoryServices)
	service := NewServiceInstitution(repo, serviceInstitutionService, serviceInstitutionClient, serviceDirectoryTree, services)
	var institution dto.CreateInstitutionRes
	t.Run("Create institution service", func(t *testing.T) {
		req := &dto.CreateInstitutionReq{
			InstitutionName: "Test",
			Description:     "Test is test",
			Clients:         []int32{1, 2},
			Services:        []int32{1, 2},
		}
		res, err := service.Create(req)
		institution = *res
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
	t.Run("Get the all institution", func(t *testing.T) {
		res, err := service.List()
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})
	t.Run("Update institution service", func(t *testing.T) {
		req := &dto.UpdateInstitutionReq{
			InstitutionID:   institution.InstitutionID,
			InstitutionName: "Test udpadte",
			Description:     "Test is test update",
			Services:        []int32{1},
		}
		res, err := service.Update(req)
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("Get  institution", func(t *testing.T) {
		res, err := service.Get(dto.GetInstitutionReq{Id: institution.InstitutionID})
		assert.NotNil(t, res)
		assert.Nil(t, err)
	})

	t.Run("Delete institution ", func(t *testing.T) {
		res, err := service.Delete(dto.GetInstitutionReq{Id: 2})
		assert.NotNil(t, res)
		assert.Nil(t, err)

	})

}
