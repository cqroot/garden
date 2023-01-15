package server

import (
	"fmt"

	"github.com/cqroot/garden/internal/app"
	"github.com/cqroot/garden/internal/databases"
	"github.com/cqroot/garden/internal/models"
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
		fmt.Sprintf("%s:%d", app.Config().BindIp(), app.Config().BindPort()),
	)

	return err
}
