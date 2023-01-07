package models

import "github.com/cqroot/todoapp/internal/databases"

func AutoMigrate() error {
	err := databases.DB().AutoMigrate(&Todo{})
	return err
}
