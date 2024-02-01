package routes

import (
	"movierental/internal/app/handlers"
	"movierental/internal/app/repository"
	"movierental/internal/app/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	repository := repository.NewRepository()
	service := service.NewService(repository)
	handlers := handlers.NewHandler(service)

	group := engine.Group("/movierental")
	{
		group.GET("/hello", handlers.GetEndPoint)
		group.GET("/movies", handlers.GetMoviesEndPoint)

	}

}
