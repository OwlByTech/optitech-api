package service

import (
	"context"
	dto "optitech/internal/dto/client"
	"optitech/internal/repository"
	sq "optitech/internal/sqlc"
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

	repoReq := sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  req.Password,
	}

	r, err := repository.Queries.CreateClient(ctx, repoReq)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
