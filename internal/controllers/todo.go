package controllers

import (
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

	c.JSON(200, todos)
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

	c.JSON(200, gin.H{
		"todo": todo,
	})
}

func GetTodo(c *gin.Context) {
	todo, err := models.GetTodo(c.GetInt("id"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}

	c.JSON(200, todo)
}
