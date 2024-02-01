package service

import (
	"movierental/internal/app/dto"
	"movierental/internal/app/repository"
)

type Service interface {
	GetEndPoint() (string, error)
	GetMoviesEndPoint() ([]dto.Movie, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetEndPoint() (string, error) {
	return s.repo.GetEndPoint()
}

func (s *service) GetMoviesEndPoint() ([]dto.Movie, error) {
	return s.repo.GetMoviesEndPoint()
}
