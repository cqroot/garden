package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/model"
)

func (c Controller) GetTasks(ctx *gin.Context) {
	dueStart, err := strconv.ParseInt(
		ctx.DefaultQuery("start", "0"),
		10, 64)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	dueEnd, err := strconv.ParseInt(
		ctx.DefaultQuery("end", "0"),
		10, 64)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	tasks, err := c.model.GetTasks(dueStart, dueEnd)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c Controller) PutTask(ctx *gin.Context) {
	var err error
	var task model.Task

	err = ctx.BindJSON(&task)
	if err != nil {
		c.middleware.AbortWithErrorAndBadRequestCode(ctx, err)
		return
	}

	err = c.model.PutTask(&task)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c Controller) UpdateTask(ctx *gin.Context) {
	var err error
	var task model.Task

	err = ctx.BindJSON(&task)
	if err != nil {
		c.middleware.AbortWithErrorAndBadRequestCode(ctx, err)
		return
	}

	if _, exists := ctx.Get("id"); exists {
		task.Id = ctx.GetUint("id")
	}

	_, err = c.model.GetTask(task.Id)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	err = c.model.UpdateTask(&task)
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c Controller) GetTask(ctx *gin.Context) {
	task, err := c.model.GetTask(ctx.GetUint("id"))
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c Controller) DeleteTask(ctx *gin.Context) {
	err := c.model.DeleteTask(ctx.GetUint("id"))
	if err != nil {
		c.middleware.AbortWithError(ctx, err)
		return
	}

	ctx.Status(http.StatusOK)
}
