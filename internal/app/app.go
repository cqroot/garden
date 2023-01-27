package app

import (
	"fmt"

	"go.uber.org/zap"
)

func Run() error {
	app, err := InitApp()
	if err != nil {
		return err
	}

	r, err := app.NewRouter()
	if err != nil {
		return err
	}

	app.logger.Debug("Listen and Server",
		zap.String("IP", app.config.BindIp()),
		zap.Int("Port", app.config.BindPort()))
	err = r.Run(
		fmt.Sprintf("%s:%d", app.config.BindIp(), app.config.BindPort()),
	)

	return err
}
