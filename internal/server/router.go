package server

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/garden/internal/middleware"
	"github.com/cqroot/garden/ui"
)

func (s Server) NewRouter() (*gin.Engine, error) {
	if s.config.LogLevel() != "Debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(s.middleware.LoggerMiddleware(func(f middleware.LogFields) {
		s.logger.Info(
			fmt.Sprintf("| %d | %15s | %15s | %-7s %s",
				f.StatusCode,
				f.Latency,
				f.ClientIp,
				f.Method,
				f.Path,
			),
		)
	}))
	router.Use(s.middleware.CorsMiddleware())
	router.Use(gin.Recovery())

	v1Group := router.Group("v1")
	{
		taskGroup := v1Group.Group("task")
		{
			taskGroup.GET("", s.controller.GetTasks)
			taskGroup.PUT("", s.controller.PutTask)
			taskGroup.PUT("/:id", s.middleware.IdValidationMiddleware, s.controller.UpdateTask)
			taskGroup.GET("/:id", s.middleware.IdValidationMiddleware, s.controller.GetTask)
			taskGroup.DELETE("/:id", s.middleware.IdValidationMiddleware, s.controller.DeleteTask)
		}

		projectGroup := v1Group.Group("project")
		{
			projectGroup.PUT("/", s.controller.UpdateProject)
			projectGroup.PUT("/:id", s.middleware.IdValidationMiddleware, s.controller.UpdateProject)
			projectGroup.DELETE("/:id", s.middleware.IdValidationMiddleware, s.controller.DeleteProject)
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
