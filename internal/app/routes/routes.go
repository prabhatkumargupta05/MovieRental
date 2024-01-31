package routes

import (
	"fmt"
	"movierental/internal/app/handlers"
	"movierental/internal/app/repository"
	"movierental/internal/app/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	repository, err := repository.NewRepository()
	if err != nil {
		fmt.Println(err)
		return
	}

	service, err := service.NewService(repository)
	if err != nil {
		fmt.Println(err)
		return
	}

	handlers, err := handlers.NewHandler(service)
	if err != nil {
		fmt.Println(err)
		return
	}

	group := engine.Group("/movierental")
	{
		group.GET("/hello", handlers.GetEndPoint)
	}

}
