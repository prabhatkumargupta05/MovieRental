package service

import (
	"movierental/internal/app/dto"
	"movierental/internal/app/repository"
)

type Service interface {
	GetEndPoint() (string, error)
	GetMoviesEndPoint() ([]dto.Movie, error)
	GetAllMovieData(string, string, string) ([]dto.Movie, error)
	GetMovieDetail(string) (dto.Movie, error)
	AddMovieToCart(string) (string, error)
    GetMoviesInCart() ([]dto.CartMovie, error)
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

func (s *service) GetAllMovieData(title string, year string, actors string) ([]dto.Movie, error) {
	return s.repo.GetAllMovieData(title, year, actors)
}

func (s *service) GetMovieDetail(imdbID string) (dto.Movie, error) {
	return s.repo.GetMovieDetail(imdbID)
}

func (s *service) AddMovieToCart(imdbID string) (string, error) {
	return s.repo.AddMovieToCart(imdbID)
}

func (s *service) GetMoviesInCart() ([]dto.CartMovie, error) {
    return s.repo.GetMoviesInCart()
}