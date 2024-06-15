package service

import (
	dto "optitech/internal/dto/client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type serviceClient struct {
	clientRepository interfaces.IClientRepository
}

func NewServiceClient(r interfaces.IClientRepository) interfaces.IClientService {
	return &serviceClient{
		clientRepository: r,
	}
}

func (s *serviceClient) Get(req dto.GetClientReq) (*dto.GetClientRes, error) {
	return s.clientRepository.GetClient(req.Id)
}

func (s *serviceClient) Create(req *dto.CreateClientReq) (*dto.CreateClientRes, error) {
	repoReq := &sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.clientRepository.CreateClient(repoReq)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *serviceClient) Update(req *dto.UpdateClientReq) (bool, error) {
	return false, nil
}

func (s *serviceClient) List() (*[]dto.GetClientRes, error) {
	return nil, nil
}

func (s *serviceClient) Delete(req dto.GetClientReq) (bool, error) {
	return false, nil
}
