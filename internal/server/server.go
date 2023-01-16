package server

import (
	"fmt"

	"github.com/cqroot/garden/internal/app"
	"github.com/cqroot/garden/internal/databases"
	"github.com/cqroot/garden/internal/models"
	"go.uber.org/zap"
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

	r, err := NewRouter()
	if err != nil {
		return err
	}

	app.Logger().Debug("Listen and Server",
		zap.String("IP", app.Config().BindIp()),
		zap.Int("Port", app.Config().BindPort()))
	err = r.Run(
		fmt.Sprintf("%s:%d", app.Config().BindIp(), app.Config().BindPort()),
	)

	return err
}
