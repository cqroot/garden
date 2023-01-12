package app

import "github.com/cqroot/todoapp/internal/configs"

func Bootstrap() {
	InitLogger(configs.LogLevel(), false)
}
