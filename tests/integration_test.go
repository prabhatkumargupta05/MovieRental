package tests

import (
	"movierental/configs"
	"movierental/database"
	"movierental/internal/app/dto"
	"movierental/internal/app/handlers"
	"movierental/internal/app/repository"
	"movierental/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func initAndRespond(requestUrl string, t *testing.T) (int, []dto.Movie, error) {
	engine := gin.Default()
	config := configs.Config{}
	configs.GetConfigs(&config)
	db := database.CreateConnection(config.Database)
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handlers.NewHandler(service)

	group := engine.Group("/movierental")
	{
		group.GET("/movies", handlers.GetAllMovieData)
	}

	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var response []dto.Movie
	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	return responseRecorder.Code, response, err
}

func TestShouldGiveListOfAllMovies(t *testing.T) {
	responseCode, respondMovies, err := initAndRespond("/movierental/movies", t)
	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, responseCode, http.StatusOK)
	assert.Equal(t, 6, len(respondMovies))

	responseCode, respondMovies, err = initAndRespond("/movierental/movies?title=batman", t)
	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, responseCode, http.StatusOK)
	assert.Equal(t, 3, len(respondMovies))

	responseCode, respondMovies, err = initAndRespond("/movierental/movies?title=batman&actors=michael", t)
	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, responseCode, http.StatusOK)
	assert.Equal(t, 2, len(respondMovies))
}
