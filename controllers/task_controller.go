package controllers

import (
	"Task-Api/database"
	"Task-Api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.Completed = input.Completed

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
