package service

import (
	cfg "optitech/internal/config"
	dto "optitech/internal/dto/client"
	dto_mailing "optitech/internal/dto/mailing"
	"optitech/internal/interfaces"
	"optitech/internal/security"
	"optitech/internal/service/mailing"
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
	hash, err := security.BcryptHashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	repoReq := &sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  hash,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	r, err := s.clientRepository.CreateClient(repoReq)
	if err != nil {
		return nil, err
	}

	client := &dto.ClientToken{
		ID: r.Id,
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &dto.CreateClientRes{
		Token: token,
	}, nil
}

func (s *serviceClient) Update(req *dto.UpdateClientReq) (bool, error) {
	repoReq := &sq.UpdateClientByIdParams{
		ClientID:  req.ClientID,
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if req.Email != "" {
		repoReq.Email = req.Email
	}

	if req.Password != "" {
		repoReq.Password = req.Password
	}

	if req.GivenName != "" {
		repoReq.GivenName = req.GivenName
	}

	if req.Surname != "" {
		repoReq.Surname = req.Surname
	}

	err := s.clientRepository.UpdateClient(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceClient) List() (*[]dto.GetClientRes, error) {
	repoRes, err := s.clientRepository.ListClient()
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (s *serviceClient) Delete(req dto.GetClientReq) (bool, error) {
	repoReq := &sq.DeleteClientByIdParams{
		ClientID:  req.Id,
		DeletedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	if err := s.clientRepository.DeleteClient(repoReq); err != nil {
		return false, pgtype.ErrScanTargetTypeChanged
	}

	return true, nil
}

func (s *serviceClient) Login(req *dto.LoginClientReq) (*dto.LoginClientRes, error) {
	res, err := s.clientRepository.LoginClient(req.Email)
	if err != nil {
		return nil, err
	}
	if security.BcryptCheckPasswordHash(req.Password, res.Password) != nil {
		return nil, err
	}

	client := &dto.ClientToken{
		ID: res.ClientID,
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &dto.LoginClientRes{
		Token: token,
	}, nil

}

func (s *serviceClient) ResetPassword(req dto.ResetPasswordReq) (bool, error) {
	res, err := s.clientRepository.LoginClient(req.Email)
	if err != nil {
		return false, err
	}
	client := &dto.ClientTokenResetPassword{
		ID:  res.ClientID,
		Exp: time.Now().Add(time.Hour / 2).Unix(),
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecretPassword)
	if err != nil {
		return false, err
	}

	if err := mailing.SendResetPassword(&dto_mailing.ResetPasswordMailingReq{
		Email:   res.Email,
		Subject: "Owlbytech Restablecer contrasena",
		Link:    cfg.Env.WebUrl + "/change-password?token=" + token,
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (s *serviceClient) ResetPasswordToken(req *dto.ResetPasswordTokenReq) (bool, error) {
	_, payload, err := security.JWTGetPayload(req.Token, cfg.Env.JWTSecretPassword)
	if err != nil {
		return false, err
	}
	client, err := s.Get(dto.GetClientReq{Id: int32(payload.ID)})
	if err != nil {
		return false, err
	}
	hash, err := security.BcryptHashPassword(req.Password)
	if err != nil {
		return false, err
	}
	res, err := s.Update(&dto.UpdateClientReq{
		ClientID:  client.ClientID,
		Password:  hash,
		Email:     client.Email,
		GivenName: client.GivenName,
		Surname:   client.Surname,
	})
	if err != nil {
		return false, err
	}

	return res, nil
}
func (s *serviceClient) ValidateResetPasswordToken(req dto.ValidateResetPasswordTokenReq) (bool, error) {
	_, err := security.JWTVerify(req.Token, cfg.Env.JWTSecretPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
