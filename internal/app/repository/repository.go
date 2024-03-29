package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"movierental/internal/app/dto"
	"net/http"
	"os"
)

type Repository interface {
	GetEndPoint() (string, error)
	GetMoviesEndPoint() ([]dto.Movie, error)
	GetAllMovieData(string, string, string) ([]dto.Movie, error)
	GetMovieDetail(string) (dto.Movie, error)
	AddMovieToCart(string) (string, error)
	GetMoviesInCart() ([]dto.CartMovie, error)
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

func (repo repository) GetAllMovieData(title string, year string, actors string) ([]dto.Movie, error) {
	var rows *sql.Rows
	var err error
	if len(title) == 0 && len(year) == 0 && len(actors) == 0 {
		rows, err = repo.Db.Query("SELECT * FROM movie")
	}
	if len(title) > 0 || len(year) > 0 || len(actors) > 0 {
		if len(title) == 0 {
			title = "%"
		}
		if len(actors) == 0 {
			actors = "%"
		}
		if len(year) == 0 {
			year = "%"
		}

		query := "SELECT * FROM movie WHERE title ILIKE $1 AND actors ILIKE $2 AND year ILIKE $3;"
		rows, err = repo.Db.Query(query, "%"+title+"%", "%"+actors+"%", "%"+year+"%")
	}
	if err != nil {
		fmt.Println("DB query failed : ", err)
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

func (repo repository) GetMovieDetail(imdbID string) (dto.Movie, error) {
	var rows *sql.Rows

	query := "SELECT * FROM movie WHERE imdb_id=$1"
	rows, err := repo.Db.Query(query, imdbID)
	if err != nil {
		fmt.Println("DB query failed : ", err)
		return dto.Movie{}, err
	}
	defer rows.Close()
	var movie dto.Movie
	for rows.Next() {
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Rated, &movie.Released, &movie.Runtime, &movie.Genre, &movie.Director, &movie.Writer, &movie.Actors, &movie.Language, &movie.Country, &movie.Awards, &movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes, &movie.ImdbID, &movie.Type, &movie.Dvd, &movie.BoxOffice, &movie.Production, &movie.Website, &movie.Response)
		if err != nil {
			return movie, errors.New("record not found")
		}
	}
	if movie.ImdbID != "" {
		return movie, nil
	}
	return movie, errors.New("record not found")
}

func (repo repository) AddMovieToCart(imdbID string) (string, error) {
	query := "INSERT INTO cart (imdb_id) VALUES ($1);"

	_, err := repo.Db.Exec(query, imdbID)
	if err != nil {
		fmt.Println("DB query failed : ", err)
		return "", err
	}
	return fmt.Sprintf("Movie (%s) added to cart", imdbID), nil
}

func (repo repository) GetMoviesInCart() ([]dto.CartMovie, error) {
	var rows *sql.Rows

	rows, err := repo.Db.Query("SELECT * FROM cart")
	if err != nil {
		fmt.Println("DB query failed : ", err)
		return nil, err
	}
	defer rows.Close()

	var cartMovies []dto.CartMovie
	for rows.Next() {

		var cartMovie dto.CartMovie
		err = rows.Scan(&cartMovie.ID, &cartMovie.ImdbID)
		if err != nil {
			return nil, err
		}
		cartMovies = append(cartMovies, cartMovie)
	}
	return cartMovies, nil
}
