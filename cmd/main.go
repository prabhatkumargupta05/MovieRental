package main

import (
	"fmt"
	"movierental/internal/app/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to Movie Rental System...")

	engine := gin.Default()
	routes.RegisterRoutes(engine)
	server := http.Server{
		Addr:                         ":8080",
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
