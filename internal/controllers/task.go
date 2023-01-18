package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/middlewares"
	"github.com/cqroot/garden/internal/models"
)

func GetTasks(c *gin.Context) {
	tasks, err := models.GetTasks()
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func PutTask(c *gin.Context) {
	var err error
	var task models.Task

	err = c.BindJSON(&task)
	if err != nil {
		middlewares.AbortWithErrorAndBadRequestCode(c, err)
		return
	}

	err = models.PutTask(&task)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var err error
	var task models.Task

	err = c.BindJSON(&task)
	if err != nil {
		middlewares.AbortWithErrorAndBadRequestCode(c, err)
		return
	}

	if _, exists := c.Get("id"); exists {
		task.Id = c.GetUint("id")
	}

	_, err = models.GetTasks()
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	err = models.UpdateTask(&task)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTask(c *gin.Context) {
	task, err := models.GetTask(c.GetUint("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	err := models.DeleteTask(c.GetUint("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
