package repository

import (
	"context"
	dto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
	"strconv"
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

func (r *repositoryFormat) CreateFormat(arg *sq.CreateFormatParams) (*dto.CreateFormatRes, error) {
	ctx := context.Background()

	res, err := r.formatRepository.CreateFormat(ctx, *arg)

	if err != nil {
		return nil, err
	}

	return &dto.CreateFormatRes{
		Id:          res.FormatID,
		AsesorId:    strconv.Itoa(int(res.AsesorID)),
		FormatName:  res.FormatName,
		Description: res.Description,
		Extension:   res.Description,
		Version:     res.Version,
	}, nil
}
