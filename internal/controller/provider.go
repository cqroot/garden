package controller

import (
	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/logger"
	"github.com/cqroot/garden/internal/middleware"
	"github.com/cqroot/garden/internal/model"
)

type Controller struct {
	config     *config.Config
	logger     *logger.Logger
	model      *model.Model
	middleware *middleware.Middleware
}

func New(config *config.Config, logger *logger.Logger, model *model.Model, middleware *middleware.Middleware) *Controller {
	return &Controller{
		config:     config,
		logger:     logger,
		model:      model,
		middleware: middleware,
	}
}
