// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/controller"
	"github.com/cqroot/garden/internal/database"
	"github.com/cqroot/garden/internal/logger"
	"github.com/cqroot/garden/internal/middleware"
	"github.com/cqroot/garden/internal/model"
)

// Injectors from wire.go:

func InitApp() (*App, error) {
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	loggerLogger := logger.New(configConfig)
	databaseDatabase, err := database.New(configConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	modelModel, err := model.New(configConfig, loggerLogger, databaseDatabase)
	if err != nil {
		return nil, err
	}
	middlewareMiddleware := middleware.New(configConfig, loggerLogger, modelModel)
	controllerController := controller.New(configConfig, loggerLogger, modelModel, middlewareMiddleware)
	app := New(configConfig, loggerLogger, middlewareMiddleware, controllerController)
	return app, nil
}
