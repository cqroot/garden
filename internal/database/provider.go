package database

import (
	"path/filepath"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/logger"
)

type Database struct {
	config *config.Config
	logger *logger.Logger
	db     *gorm.DB
}

func New(config *config.Config, logger *logger.Logger) (*Database, error) {
	dsn := filepath.Join(config.DataPath(), "garden.db")
	logger.Debug("Open database", zap.String("DB", dsn))

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		config: config,
		logger: logger,
		db:     db,
	}, nil
}

func (d Database) DB() *gorm.DB {
	return d.db
}
