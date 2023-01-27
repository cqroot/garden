package app

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/middleware"
	"github.com/cqroot/garden/ui"
)

func (app App) NewRouter() (*gin.Engine, error) {
	if app.config.LogLevel() != "Debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(app.middleware.LoggerMiddleware(func(f middleware.LogFields) {
		app.logger.Info(
			fmt.Sprintf("| %d | %15s | %15s | %-7s %s",
				f.StatusCode,
				f.Latency,
				f.ClientIp,
				f.Method,
				f.Path,
			),
		)
	}))
	router.Use(app.middleware.CorsMiddleware())
	router.Use(gin.Recovery())

	v1Group := router.Group("v1")
	{
		taskGroup := v1Group.Group("task")
		{
			taskGroup.GET("", app.controller.GetTasks)
			taskGroup.PUT("", app.controller.PutTask)
			taskGroup.PUT("/:id", app.middleware.IdValidationMiddleware, app.controller.UpdateTask)
			taskGroup.GET("/:id", app.middleware.IdValidationMiddleware, app.controller.GetTask)
			taskGroup.DELETE("/:id", app.middleware.IdValidationMiddleware, app.controller.DeleteTask)
		}

		projectGroup := v1Group.Group("project")
		{
			projectGroup.PUT("/", app.controller.UpdateProject)
			projectGroup.PUT("/:id", app.middleware.IdValidationMiddleware, app.controller.UpdateProject)
			projectGroup.DELETE("/:id", app.middleware.IdValidationMiddleware, app.controller.DeleteProject)
		}
	}

	fDist, err := fs.Sub(ui.FEmbedUi, "dist")
	if err != nil {
		return nil, err
	}
	router.StaticFS("/ui", http.FS(fDist))
	router.NoRoute(func(c *gin.Context) {
		c.FileFromFS("index.html", http.FS(fDist))
	})

	return router, nil
}
