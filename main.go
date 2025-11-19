package main

import (
	"Task-Api/database"
	"Task-Api/routes"

	_ "Task-Api/docs" // ต้องมีเพื่อให้ Swagger รู้จัก docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	controllers "Task-Api/controllers"
)

// @title Task API
// @version 1.0
// @description Simple Task API with Gin, GORM, PostgreSQL
// @host localhost:8000
// @BasePath /
func main() {
	database.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/tasks ", controllers.GetTasks)
	r.Run(":8000")
}
