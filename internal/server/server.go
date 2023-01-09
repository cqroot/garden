package server

import (
	"fmt"

	"github.com/cqroot/todoapp/internal/configs"
	"github.com/cqroot/todoapp/internal/databases"
	"github.com/cqroot/todoapp/internal/models"
	"github.com/cqroot/todoapp/internal/utils"
)

func Run() error {
	var err error

	utils.InitLogger(configs.LogLevel(), false)

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
