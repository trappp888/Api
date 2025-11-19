package routes

import (
	"Task-Api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
}
