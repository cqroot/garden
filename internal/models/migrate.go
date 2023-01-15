package models

import "github.com/cqroot/garden/internal/databases"

func AutoMigrate() error {
	var err error

	err = databases.DB().AutoMigrate(&Task{})
	if err != nil {
		return err
	}

	err = databases.DB().AutoMigrate(&Project{})
	if err != nil {
		return err
	}

	return nil
}
