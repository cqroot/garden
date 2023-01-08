package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/todoapp/internal/controllers"
	"github.com/cqroot/todoapp/internal/middlewares"
	"github.com/cqroot/todoapp/internal/utils"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(middlewares.CustomLogger(func(f middlewares.LogFields) {
		utils.Logger().Info(
			fmt.Sprintf("| %d | %15s | %15s | %-7s %s",
				f.StatusCode,
				f.Latency,
				f.ClientIp,
				f.Method,
				f.Path,
			),
		)
	}))
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
