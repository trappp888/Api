package main

import (
	"Task-Api/database"
	"Task-Api/routes"

	"github.com/gin-gonic/gin"

	_ "Task-Api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Task API
// @version 1.0
// @description Simple Task API with Gin, Gorm, PostgreSQL

// @host localhost:8080
// @BasePath /

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
