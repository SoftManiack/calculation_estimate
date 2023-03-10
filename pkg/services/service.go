package service

import (
	dto "calculation-estimate"
)

type Inquirer interface {
	GetInquirer() (dto.Inquirer, error)
	–°alculationEstimate(inquirer dto.Inquirer) ([]dto.Estimate, error)
}

type Service struct {
	Inquirer
}

func NewService() *Service {
	return &Service{
		Inquirer: NewInquirerService(),
	}
}
