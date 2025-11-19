package controllers

import (
	"Task-Api/database"
	"Task-Api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTasks godoc
// @Summary Get all tasks
// @Description Get list of tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary Create a new task
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Task Data"
// @Success 200 {object} models.Task
// @Router /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update a task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body models.Task true "Task Data"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [put]
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

// DeleteTask godoc
// @Summary Delete a task
// @Tags tasks
// @Param id path int true "Task ID"
// @Success 200 {object} map[string]string
// @Router /tasks/{id} [delete]
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
