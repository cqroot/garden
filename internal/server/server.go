package server

import (
	"github.com/cqroot/todoapp/internal/databases"
	"github.com/cqroot/todoapp/internal/models"
)

func Run() error {
	var err error

	err = databases.InitDatabase()
	if err != nil {
		return err
	}

	err = models.AutoMigrate()
	if err != nil {
		return err
	}

	r := NewRouter()
	err = r.Run("127.0.0.1:8000")

	return err
}
