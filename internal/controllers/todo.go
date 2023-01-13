package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/todoapp/internal/middlewares"
	"github.com/cqroot/todoapp/internal/models"
)

func GetTodos(c *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	var err error
	var todo models.Todo

	err = c.BindJSON(&todo)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	if _, exists := c.Get("id"); exists {
		todo.Id = c.GetInt("id")
	}

	err = models.UpdateTodo(&todo)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func GetTodo(c *gin.Context) {
	todo, err := models.GetTodo(c.GetInt("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	err := models.DeleteTodo(c.GetInt("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func MarkTodoDone(c *gin.Context) {
	err := models.MarkTodoDone(c.GetInt("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
