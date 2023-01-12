package server

import (
	"fmt"

	"github.com/cqroot/todoapp/internal/configs"
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
	err = r.Run(
		fmt.Sprintf("%s:%d", configs.BindIp(), configs.BindPort()),
	)

	return err
}
