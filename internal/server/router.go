package server

import (
	"github.com/gin-gonic/gin"

	"github.com/cqroot/todoapp/internal/controllers"
	"github.com/cqroot/todoapp/internal/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1Group := router.Group("v1")
	{
		todoGroup := v1Group.Group("todo")
		{
			todoGroup.GET("", controllers.GetTodos)
			todoGroup.PUT("", controllers.UpdateTodo)
			todoGroup.PUT("/:id", middlewares.ValidateIdMiddleware, controllers.UpdateTodo)
			todoGroup.GET("/:id", middlewares.ValidateIdMiddleware, controllers.GetTodo)
		}
	}

	return router

}
