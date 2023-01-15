package databases

import (
	"github.com/adrg/xdg"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/cqroot/garden/internal/app"
)

var db *gorm.DB

func InitDatabase() error {
	dsn, err := xdg.DataFile("garden")
	if err != nil {
		return err
	}
	app.Logger().Debug("Open database", zap.String("db", dsn))

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func DB() *gorm.DB {
	return db
}
