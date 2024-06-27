package repository

import (
	"context"
	dto "optitech/internal/dto/asesor"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryAsesor struct {
	asesorRepository *sq.Queries
}

func NewRepositoryAsesor(q *sq.Queries) interfaces.IAsesorRepository {
	return &repositoryAsesor{
		asesorRepository: q,
	}
}

func (r *repositoryAsesor) GetAsesor(asesorID int32) (*dto.GetAsesorRes, error) {
	ctx := context.Background()

	repoRes, err := r.asesorRepository.GetAsesor(ctx, (asesorID))

	if err != nil {
		return nil, err
	}

	return &dto.GetAsesorRes{
		Id:    repoRes.AsesorID,
		About: repoRes.About,
	}, nil
}

func (r *repositoryAsesor) CreateAsesor(arg *sq.CreateAsesorParams) (*dto.CreateAsesorRes, error) {
	ctx := context.Background()

	res, err := r.asesorRepository.CreateAsesor(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateAsesorRes{
		Id:    res.AsesorID,
		About: res.About,
	}, nil
}

func (r *repositoryAsesor) UpdateAsesor(arg *sq.UpdateAsesorByIdParams) error {
	ctx := context.Background()
	return r.asesorRepository.UpdateAsesorById(ctx, *arg)
}

func (r *repositoryAsesor) ListAsesor() (*[]dto.GetAsesorRes, error) {
	ctx := context.Background()
	repoRes, err := r.asesorRepository.ListAsesors(ctx)

	if err != nil {
		return nil, err
	}

	asesors := make([]dto.GetAsesorRes, len(repoRes))
	for i, inst := range repoRes {
		asesors[i] = dto.GetAsesorRes{
			Id:    inst.AsesorID,
			About: inst.About,
		}
	}
	return &asesors, nil
}

func (r *repositoryAsesor) DeleteAsesor(arg *sq.DeleteAsesorByIdParams) error {
	ctx := context.Background()
	return r.asesorRepository.DeleteAsesorById(ctx, *arg)
}
