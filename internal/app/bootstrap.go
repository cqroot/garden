package app

import "github.com/cqroot/garden/internal/configs"

func Bootstrap() {
	InitLogger(configs.LogLevel(), false)
}
