package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/model"
)

func (c Controller) UpdateProject(ctx *gin.Context) {
	var err error
	var project model.Project

	err = ctx.BindJSON(&project)
	if err != nil {
		c.middleware.AbortWithErrorAndBadRequestCode(ctx, err)
		return
	}
	if _, exists := ctx.Get("id"); exists {
		project.Id = ctx.GetInt("id")
	}

	err = c.model.UpdateProject(&project)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (c Controller) DeleteProject(ctx *gin.Context) {
	err := c.model.DeleteProject(ctx.GetInt("id"))
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
