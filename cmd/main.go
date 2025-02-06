// @title Buildings API
// @version 1.0
// @description API for managing buildings.
// @host localhost:8080
// @BasePath /
package main

import (
	_ "leadgen/docs"
	"leadgen/internal/adapter/http"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := http.NewServer()
	defer server.Close()

	server.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Starting server on :8080")
	server.Run(":8080")
}
