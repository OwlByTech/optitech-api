package service

import (
	dto_i "optitech/internal/dto/institution"
	"optitech/internal/interfaces"
)

type serviceDirectoryInstitution struct {
	DirectoryInstitutionRepository interfaces.IDirectoryInstitutionRepository
}

func (s *serviceDirectoryInstitution) ListInstitutionByDirectoryId(directoryId int32) (*[]dto_i.GetInstitutionRes, error) {
	repoRes, err := s.DirectoryInstitutionRepository.ListInstitutionByDirectoryId(directoryId)
	if err != nil {
		return nil, err
	}

	institutions := make([]dto_i.GetInstitutionRes, len(*repoRes))

	for i, inst := range *repoRes {
		institutions[i] =
			dto_i.GetInstitutionRes{
				Id:              inst.Institution.InstitutionID,
				InstitutionName: inst.Institution.InstitutionName,
				Description:     inst.Institution.Description,
			}
	}

	return &institutions, nil

}
