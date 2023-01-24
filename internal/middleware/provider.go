package middleware

import (
	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/logger"
	"github.com/cqroot/garden/internal/model"
)

type Middleware struct {
	config *config.Config
	logger *logger.Logger
	model  *model.Model
}

func New(config *config.Config, logger *logger.Logger, model *model.Model) *Middleware {
	return &Middleware{
		config: config,
		logger: logger,
		model:  model,
	}
}
