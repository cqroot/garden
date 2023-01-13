package models

import (
	"github.com/cqroot/todoapp/internal/databases"
	"gorm.io/gorm/clause"
)

type Todo struct {
	Id    int    `json:"id"    gorm:"unique;AUTO_INCREMENT"`
	Title string `json:"title" gorm:"not null"`
	Note  string `json:"note"`
	Done  bool   `json:"done" gorm:"default:false"`
}

func GetTodos() (*[]Todo, error) {
	var todos []Todo

	err := databases.DB().Find(&todos).Error
	return &todos, err
}

func UpdateTodo(todo *Todo) error {
	// err := databases.DB().Create(todo).Error
	err := databases.DB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(todo).Error

	return err
}

func GetTodo(id int) (*Todo, error) {
	var todo Todo

	err := databases.DB().First(&todo, id).Error
	return &todo, err
}

func DeleteTodo(id int) error {
	return databases.DB().Delete(&Todo{}, id).Error
}

func MarkTodoDone(id int) error {
	return databases.DB().
		Model(&Todo{Id: id}).
		Update("done", true).Error
}
