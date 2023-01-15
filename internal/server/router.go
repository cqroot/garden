package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/app"
	"github.com/cqroot/garden/internal/controllers"
	"github.com/cqroot/garden/internal/middlewares"
)

func NewRouter() *gin.Engine {
	if app.Config().LogLevel() != "Debug" {
		gin.SetMode(gin.ReleaseMode)
	}

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
		taskGroup := v1Group.Group("task")
		{
			taskGroup.GET("", controllers.GetTasks)
			taskGroup.PUT("", controllers.UpdateTask)
			taskGroup.PUT("/:id", middlewares.IdValidationMiddleware, controllers.UpdateTask)
			taskGroup.GET("/:id", middlewares.IdValidationMiddleware, controllers.GetTask)
			taskGroup.DELETE("/:id", middlewares.IdValidationMiddleware, controllers.DeleteTask)
			taskGroup.PUT("/status/:id", middlewares.IdValidationMiddleware, controllers.MarkTaskDone)
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
