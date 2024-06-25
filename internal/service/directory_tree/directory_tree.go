package service

import (
	dto "optitech/internal/dto/directory_tree"
	"optitech/internal/interfaces"
)

type serviceDirectoryTree struct {
	directoryTreeRepository interfaces.IDirectoryRepositoy
}

func NewServicDirectory(r interfaces.IDirectoryRepositoy) interfaces.IDirectoryService {
	return &serviceDirectoryTree{
		directoryTreeRepository: r,
	}
}

func (s *serviceDirectoryTree) Get(req dto.GetDirectoryTreeReq) (*dto.GetDirectoryTreeRes, error) {
	return s.directoryTreeRepository.GetDirectroy(req.Id)
}

func (s *serviceDirectoryTree) Create(req dto.CreateDirectoryTreeReq) (*dto.CreateDirectoryTreeRes, error) {
	return nil, nil
}
