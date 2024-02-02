package main

import (
	"fmt"
	"movierental/configs"
	"movierental/internal/app/routes"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Movie Rental System...")

	config := configs.Config{}
	configs.GetConfigs(&config)
	
	engine := gin.Default()
	routes.RegisterRoutes(engine, config.Database)

	server := http.Server{
		Addr:                         ":" + strconv.Itoa(config.Server.Port),
		Handler:                      engine,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		return
	}

}
