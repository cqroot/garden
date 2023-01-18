package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/middlewares"
	"github.com/cqroot/garden/internal/models"
)

func UpdateProject(c *gin.Context) {
	var err error
	var project models.Project

	err = c.BindJSON(&project)
	if err != nil {
		middlewares.AbortWithErrorAndBadRequestCode(c, err)
		return
	}
	if _, exists := c.Get("id"); exists {
		project.Id = c.GetInt("id")
	}

	err = models.UpdateProject(&project)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
	err := models.DeleteProject(c.GetInt("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
