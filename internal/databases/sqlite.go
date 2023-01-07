package databases

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() error {
	var err error

	dialector := sqlite.Open("todo.db")
	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func DB() *gorm.DB {
	return db
}
