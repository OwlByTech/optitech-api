package service

import (
	"context"
	"fmt"
	"io"
	dto "optitech/internal/dto/institution"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func GetInstitutionService(req dto.GetInstitutionReq) (*dto.GetInstitutionRes, error) {
	ctx := context.Background()

	repoRes, err := repository.Queries.GetInstitution(ctx, req.InstitutionID)

	if err != nil {
		return nil, err
	}

	return &dto.GetInstitutionRes{
		InstitutionID:   repoRes.InstitutionID,
		InstitutionName: repoRes.InstitutionName,
		Description:     repoRes.Description,
	}, nil
}

func ListInstitutionsService() ([]*dto.GetInstitutionRes, error) {
	ctx := context.Background()
	repoRes, err := repository.Queries.ListInstitutions(ctx)
	if err != nil {
		return nil, err
	}

	institutions := make([]*dto.GetInstitutionRes, len(repoRes))
	for i, inst := range repoRes {
		services := []string{}
		institutions[i] = &dto.GetInstitutionRes{
			InstitutionID:   inst.InstitutionID,
			Description:     inst.Description,
			InstitutionName: inst.InstitutionName,
			Logo:            inst.Logo.String,
			Services:        services,
		}
	}

	return institutions, nil
}

func CreateInstitutionService(req dto.CreateInstitutionReq) (*sq.Institution, error) {
	ctx := context.Background()
	repoReq := sq.CreateInstitutionParams{
		InstitutionName: req.InstitutionName,
		Description:     req.Description,
		CreatedAt:       pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.LogoFile != nil {
		nameFile := req.InstitutionName + "_" + req.LogoFile.Filename
		multipart, err := req.LogoFile.Open()
		if err != nil {
			return nil, err
		}
		defer multipart.Close()
		savePath := fmt.Sprintf("./uploads/%s", nameFile)

		outFile, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer outFile.Close()
		if _, err = io.Copy(outFile, multipart); err != nil {
			return nil, err
		}
		repoReq.Logo = pgtype.Text{String: nameFile, Valid: true}
	}

	if req.AsesorID < 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}

	r, err := repository.Queries.CreateInstitution(ctx, repoReq)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func UpdateInstitutionService(req dto.UpdateInstitutionReq) (bool, error) {
	ctx := context.Background()
	repoReq := sq.UpdateInstitutionByIdParams{
		InstitutionID: req.InstitutionID,
	}
	if req.AsesorID < 0 {
		repoReq.AsesorID = pgtype.Int4{Int32: req.AsesorID, Valid: true}
	}
	if req.InstitutionName != "" {
		repoReq.InstitutionName = req.InstitutionName
	}
	if req.Description != "" {
		repoReq.Description = req.Description
	}

	if req.LogoFile != nil {
		nameFile := req.InstitutionName + "_" + req.LogoFile.Filename
		multipart, err := req.LogoFile.Open()
		if err != nil {
			return false, err
		}
		defer multipart.Close()
		savePath := fmt.Sprintf("./uploads/%s", nameFile)

		outFile, err := os.Create(savePath)
		if err != nil {
			return false, err
		}
		defer outFile.Close()
		if _, err = io.Copy(outFile, multipart); err != nil {
			return false, err
		}
		repoReq.Logo = pgtype.Text{String: nameFile, Valid: true}
	}

	err := repository.Queries.UpdateInstitutionById(ctx, repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteInstitutionService(req dto.GetInstitutionReq) (bool, error) {
	ctx := context.Background()
	repoReq := sq.DeleteInstitutionByIdParams{
		InstitutionID: req.InstitutionID,
		DeletedAt:     pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	err := repository.Queries.DeleteInstitutionById(ctx, repoReq)

	if err != nil {
		return false, err
	}

	return true, err
}
