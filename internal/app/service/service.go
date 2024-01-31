package service

import "movierental/internal/app/repository"

type Service interface {
	GetEndPoint() (string, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) (Service, error) {
	return &service{repo: repo}, nil
}

func (s *service) GetEndPoint() (string, error) {
	return s.repo.GetEndPoint()
}
