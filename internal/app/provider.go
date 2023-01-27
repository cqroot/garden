package app

import (
	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/controller"
	"github.com/cqroot/garden/internal/logger"
	"github.com/cqroot/garden/internal/middleware"
)

type App struct {
	config     *config.Config
	logger     *logger.Logger
	middleware *middleware.Middleware
	controller *controller.Controller
}

func New(config *config.Config, logger *logger.Logger, middleware *middleware.Middleware, controller *controller.Controller) *App {
	return &App{
		config:     config,
		logger:     logger,
		middleware: middleware,
		controller: controller,
	}
}
