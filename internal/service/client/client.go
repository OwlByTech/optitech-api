package service

import (
	"context"
	dto "optitech/internal/dto/client"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func GetClientService(req dto.GetClientReq) (*dto.GetClientRes, error) {
	ctx := context.Background()

	repoRes, err := repository.Queries.GetClient(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRes{
		Id:        repoRes.ClientID,
		GivenName: repoRes.GivenName,
		Surname:   repoRes.Surname,
		Email:     repoRes.Email,
	}, nil
}

func CreateClientService(req dto.CreateClientReq) (*sq.Client, error) {
	ctx := context.Background()

	// Time
	createdAt := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}

	repoReq := sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: createdAt,
	}

	r, err := repository.Queries.CreateClient(ctx, repoReq)

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func UpdateClientService(req dto.UpdateClientReq) error {
	ctx := context.Background()

	// Time
	updatedAt := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}

	repoReq := sq.UpdateClientByIdParams{
		ClientID:  req.ClientID,
		GivenName: req.GivenName,
		Password:  req.Password,
		Surname:   req.Surname,
		Email:     req.Email,
		UpdatedAt: updatedAt,
	}

	err := repository.Queries.UpdateClientById(ctx, repoReq)
	if err != nil {
		return err
	}

	return nil
}
