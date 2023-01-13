package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/todoapp/internal/app"
	"github.com/cqroot/todoapp/internal/controllers"
	"github.com/cqroot/todoapp/internal/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(middlewares.LoggerMiddleware(func(f middlewares.LogFields) {
		app.Logger().Info(
			fmt.Sprintf("| %d | %15s | %15s | %-7s %s",
				f.StatusCode,
				f.Latency,
				f.ClientIp,
				f.Method,
				f.Path,
			),
		)
	}))
	router.Use(middlewares.CorsMiddleware())
	router.Use(gin.Recovery())

	v1Group := router.Group("v1")
	{
		todoGroup := v1Group.Group("todo")
		{
			todoGroup.GET("", controllers.GetTodos)
			todoGroup.PUT("", controllers.UpdateTodo)
			todoGroup.PUT("/:id", middlewares.IdValidationMiddleware, controllers.UpdateTodo)
			todoGroup.GET("/:id", middlewares.IdValidationMiddleware, controllers.GetTodo)
			todoGroup.DELETE("/:id", middlewares.IdValidationMiddleware, controllers.DeleteTodo)

			todoGroup.PUT("/status/:id", middlewares.IdValidationMiddleware, controllers.MarkTodoDone)
		}
	}

	return router
}
