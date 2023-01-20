package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/middlewares"
	"github.com/cqroot/garden/internal/models"
)

func GetTasks(c *gin.Context) {
	dueStart, err := strconv.ParseInt(
		c.DefaultQuery("start", "0"),
		10, 64)
	dueEnd, err := strconv.ParseInt(
		c.DefaultQuery("end", "0"),
		10, 64)

	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	tasks, err := models.GetTasks(dueStart, dueEnd)
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

	_, err = models.GetTask(task.Id)
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
