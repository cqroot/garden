package models

import "github.com/cqroot/todoapp/internal/databases"

func AutoMigrate() error {
	var err error

	err = databases.DB().AutoMigrate(&Todo{})
	if err != nil {
		return err
	}

	err = databases.DB().AutoMigrate(&Project{})
	if err != nil {
		return err
	}

	return nil
}
