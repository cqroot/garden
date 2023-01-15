package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/app"
	"github.com/cqroot/garden/internal/controllers"
	"github.com/cqroot/garden/internal/middlewares"
)

func NewRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)

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
			todoGroup.GET("", controllers.GetTasks)
			todoGroup.PUT("", controllers.UpdateTask)
			todoGroup.PUT("/:id", middlewares.IdValidationMiddleware, controllers.UpdateTask)
			todoGroup.GET("/:id", middlewares.IdValidationMiddleware, controllers.GetTask)
			todoGroup.DELETE("/:id", middlewares.IdValidationMiddleware, controllers.DeleteTask)
			todoGroup.PUT("/status/:id", middlewares.IdValidationMiddleware, controllers.MarkTaskDone)
		}

		projectGroup := v1Group.Group("project")
		{
			projectGroup.PUT("/", controllers.UpdateProject)
			projectGroup.PUT("/:id", middlewares.IdValidationMiddleware, controllers.UpdateProject)
			projectGroup.DELETE("/:id", middlewares.IdValidationMiddleware, controllers.DeleteProject)
		}
	}

	return router
}
