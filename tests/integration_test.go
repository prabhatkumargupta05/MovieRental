package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"movierental/configs"
	"movierental/database"
	"movierental/internal/app/dto"
	"movierental/internal/app/handlers"
	"movierental/internal/app/repository"
	"movierental/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUp() *gin.Engine {
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
	return engine
}
func TestShouldGiveListOfAllMovies(t *testing.T) {
	engine := setUp()
	request, err := http.NewRequest(http.MethodGet, "/movierental/movies", nil)

	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	//what should be variable type here
	var response []dto.Movie

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, 6, len(response))

}

func TestShouldGiveListOfFilteredMovies(t *testing.T) {
	engine := setUp()
	request, err := http.NewRequest(http.MethodGet, "/movierental/movies?title=batman", nil)

	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	//what should be variable type here
	var response []dto.Movie

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	// Assert
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, 3, len(response))

}
