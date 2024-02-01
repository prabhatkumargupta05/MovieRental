package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"movierental/internal/app/dto"
	"net/http"
	"os"
)

type Repository interface {
	GetEndPoint() (string, error)
	GetMoviesEndPoint() ([]dto.Movie, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo repository) GetEndPoint() (string, error) {
	_ = repo
	return "helloworld...", nil
}

func (repo repository) GetMoviesEndPoint() ([]dto.Movie, error) {

	req, err := http.NewRequest(http.MethodGet, "https://www.omdbapi.com/?s=batman&apikey=4c36c62a", nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)

	moviesResponse := dto.MoviesResponse{}
	err = json.Unmarshal(resBody, &moviesResponse)
	if err != nil {
		fmt.Printf("client: could not unmarshal response body: %s\n", err)
		os.Exit(1)
	}
	return moviesResponse.Movies, nil
}
