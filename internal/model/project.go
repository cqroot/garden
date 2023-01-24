package model

import (
	"gorm.io/gorm/clause"
)

type Project struct {
	Id    int    `json:"id"    gorm:"unique;AUTO_INCREMENT"`
	Title string `json:"title" gorm:"not null"`
}

func (m Model) UpdateProject(project *Project) error {
	err := m.database.DB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(project).Error

	return err
}

func (m Model) DeleteProject(id int) error {
	return m.database.DB().Delete(&Project{}, id).Error
}
