package repository

import (
	"context"
	dto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type repositoryFormat struct {
	formatRepository *sq.Queries
}

func NewRepositoryFormat(q *sq.Queries) interfaces.IFormatRepository {
	return &repositoryFormat{
		formatRepository: q,
	}
}

func (r *repositoryFormat) GetFormat(formatID int32) (*dto.GetFormatRes, error) {
	ctx := context.Background()

	repoRes, err := r.formatRepository.GetFormat(ctx, (formatID))

	if err != nil {
		return nil, err
	}

	return &dto.GetFormatRes{
		Id:          repoRes.FormatID,
		AsesorId:    repoRes.AsesorID,
		Description: repoRes.Description,
		//TODO: items field is in?
		Extension: string(repoRes.Extension),
		Version:   repoRes.Version,
	}, nil
}

func (r *repositoryFormat) ListById(arg *sq.ListFormatsByIdParams) (*[]dto.GetFormatRes, error) {
	ctx := context.Background()
	repoRes, err := r.formatRepository.ListFormatsById(ctx, *arg)

	if err != nil {
		return nil, err
	}

	formats := make([]dto.GetFormatRes, len(repoRes))
	for i, forms := range repoRes {
		formats[i] = dto.GetFormatRes{
			Id:          forms.FormatID,
			AsesorId:    forms.AsesorID,
			Description: forms.Description,
			Extension:   string(forms.Extension),
			Version:     forms.Version,
		}
	}
	return &formats, nil
}

func (r *repositoryFormat) CreateFormat(arg *sq.CreateFormatParams) (*dto.CreateFormatRes, error) {
	ctx := context.Background()

	res, err := r.formatRepository.CreateFormat(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateFormatRes{
		Id:          res.FormatID,
		AsesorId:    res.AsesorID,
		Name:        res.FormatName,
		Description: res.Description,
		Extension:   res.Description,
		Version:     res.Version,
	}, nil
}

func (r *repositoryFormat) List() (*[]dto.GetFormatRes, error) {
	ctx := context.Background()
	repoRes, err := r.formatRepository.ListFormats(ctx)

	if err != nil {
		return nil, err
	}

	formats := make([]dto.GetFormatRes, len(repoRes))
	for i, forms := range repoRes {
		formats[i] = dto.GetFormatRes{
			Id:          forms.FormatID,
			AsesorId:    forms.AsesorID,
			Description: forms.Description,
			Extension:   string(forms.Extension),
			Version:     forms.Version,
		}
	}
	return &formats, nil
}

func (r *repositoryFormat) DeleteFormat(arg *sq.DeleteFormatByIdParams) error {
	ctx := context.Background()
	return r.formatRepository.DeleteFormatById(ctx, *arg)
}

func (r *repositoryFormat) UpdateFormat(arg *sq.UpdateFormatByIdParams) error {
	ctx := context.Background()
	return r.formatRepository.UpdateFormatById(ctx, *arg)
}
