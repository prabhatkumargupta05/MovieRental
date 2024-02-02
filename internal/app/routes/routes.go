package routes

import (
	"movierental/configs"
	"movierental/database"
	"movierental/internal/app/handlers"
	"movierental/internal/app/repository"
	"movierental/internal/app/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, dbConfig configs.DatabaseConfig) {
	db := database.CreateConnection(dbConfig)
	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handlers.NewHandler(service)

	group := engine.Group("/movierental")
	{
		group.GET("/hello", handlers.GetEndPoint)
		group.GET("/moviesv0", handlers.GetMoviesEndPoint)
		group.GET("/movies", handlers.GetAllMovieData)
		group.GET("/movies/:imdbID", handlers.GetMovieDetail)
	}

}
