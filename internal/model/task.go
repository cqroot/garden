package model

import (
	"gorm.io/gorm/clause"
)

type TaskStatus uint8

const (
	Todo  TaskStatus = 0
	Done  TaskStatus = 1
	Doing TaskStatus = 2
)

type Task struct {
	Id      uint   `json:"id"    gorm:"unique;AUTO_INCREMENT"`
	Title   string `json:"title" gorm:"not null"`
	Note    string `json:"note"`
	Project int    `json:"project"`
	Due     int64  `json:"due"`
	Status  uint8  `json:"done" gorm:"default:0"`
}

func (m Model) GetTasks(dueStart int64, dueEnd int64) (*[]Task, error) {
	var tasks []Task

	var err error
	if dueStart != 0 && dueEnd != 0 {
		err = m.database.DB().
			Where("due >= ?", dueStart).
			Where("due <= ?", dueEnd).
			Find(&tasks).Error
	} else if dueStart == 0 && dueEnd != 0 {
		err = m.database.DB().
			Where("due <= ?", dueEnd).
			Find(&tasks).Error
	} else if dueStart != 0 && dueEnd == 0 {
		err = m.database.DB().
			Where("due >= ?", dueStart).
			Find(&tasks).Error
	} else {
		err = m.database.DB().Find(&tasks).Error
	}
	return &tasks, err
}

func (m Model) PutTask(task *Task) error {
	return m.database.DB().Create(task).Error
}

func (m Model) UpdateTask(task *Task) error {
	err := m.database.DB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(task).Error

	return err
}

func (m Model) GetTask(id uint) (*Task, error) {
	var task Task

	err := m.database.DB().First(&task, id).Error
	return &task, err
}

func (m Model) DeleteTask(id uint) error {
	return m.database.DB().Delete(&Task{}, id).Error
}
