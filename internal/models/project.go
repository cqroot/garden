package models

import (
	"github.com/cqroot/garden/internal/databases"
	"gorm.io/gorm/clause"
)

type Project struct {
	Id    int    `json:"id"    gorm:"unique;AUTO_INCREMENT"`
	Title string `json:"title" gorm:"not null"`
}

func UpdateProject(project *Project) error {
	err := databases.DB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(project).Error

	return err
}

func DeleteProject(id int) error {
	return databases.DB().Delete(&Project{}, id).Error
}
