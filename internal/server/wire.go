//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"

	"github.com/cqroot/garden/internal/config"
	"github.com/cqroot/garden/internal/controller"
	"github.com/cqroot/garden/internal/database"
	"github.com/cqroot/garden/internal/logger"
	"github.com/cqroot/garden/internal/middleware"
	"github.com/cqroot/garden/internal/model"
)

func InitServer() (*Server, error) {
	wire.Build(
		config.New,
		logger.New,
		database.New,
		model.New,
		middleware.New,
		controller.New,
		New,
	)
	return &Server{}, nil
}
