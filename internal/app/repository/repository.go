package repository

import (
	"database/sql"
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
	GetAllMovieData() ([]dto.Movie, error)
}

type repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{Db: db}
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

func (repo repository) GetAllMovieData() ([]dto.Movie, error) {
	rows, err := repo.Db.Query("SELECT * FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []dto.Movie
	for rows.Next() {
		var movie dto.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Rated, &movie.Released, &movie.Runtime, &movie.Genre, &movie.Director, &movie.Writer, &movie.Actors, &movie.Language, &movie.Country, &movie.Awards, &movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes, &movie.ImdbID, &movie.Type, &movie.Dvd, &movie.BoxOffice, &movie.Production, &movie.Website, &movie.Response)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
