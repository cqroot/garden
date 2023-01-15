package models

import (
	"time"

	"github.com/cqroot/garden/internal/databases"
	"gorm.io/gorm/clause"
)

type TaskStatus uint8

const (
	Todo  TaskStatus = 0
	Done  TaskStatus = 1
	Doing TaskStatus = 2
)

type Task struct {
	Id      int       `json:"id"    gorm:"unique;AUTO_INCREMENT"`
	Title   string    `json:"title" gorm:"not null"`
	Note    string    `json:"note"`
	Project int       `json:"project"`
	Due     time.Time `json:"due"`
	Status  uint8     `json:"done" gorm:"default:0"`
}

func GetTasks() (*[]Task, error) {
	var tasks []Task

	err := databases.DB().Find(&tasks).Error
	return &tasks, err
}

func UpdateTask(task *Task) error {
	err := databases.DB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(task).Error

	return err
}

func GetTask(id int) (*Task, error) {
	var task Task

	err := databases.DB().First(&task, id).Error
	return &task, err
}

func DeleteTask(id int) error {
	return databases.DB().Delete(&Task{}, id).Error
}

func MarkTaskDone(id int) error {
	return databases.DB().
		Model(&Task{Id: id}).
		Update("status", Done).Error
}
