package databases

import (
	"path/filepath"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/cqroot/garden/internal/app"
)

var db *gorm.DB

func InitDatabase() error {
	var err error

	dsn := filepath.Join(app.Config().DataPath(), "garden.db")
	app.Logger().Debug("Open database", zap.String("DB", dsn))

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func DB() *gorm.DB {
	return db
}
