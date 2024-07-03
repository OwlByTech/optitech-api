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
