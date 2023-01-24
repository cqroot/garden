package model

import (
	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/database"
	"github.com/cqroot/garden/internal/logger"
)

type Model struct {
	config   *config.Config
	logger   *logger.Logger
	database *database.Database
}

func New(config *config.Config, logger *logger.Logger, database *database.Database) (*Model, error) {
	model := Model{
		config:   config,
		logger:   logger,
		database: database,
	}

	err := model.AutoMigrate()
	return &model, err
}
