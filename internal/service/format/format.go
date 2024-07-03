package service

import (
	dto "optitech/internal/dto/format"
	"optitech/internal/interfaces"
)

type serviceFormat struct {
	formatRepository interfaces.IFormatRepository
}

func NewServiceFormat(f interfaces.IFormatRepository) interfaces.IFormatService {
	return &serviceFormat{
		formatRepository: f,
	}
}

func (s *serviceFormat) Get(req dto.GetFormatReq) (*dto.GetFormatRes, error) {
	return s.formatRepository.GetFormat(req.Id)
}
