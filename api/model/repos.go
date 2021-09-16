package model

import (
	"github.com/jinzhu/gorm"
)

func FindTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Find(&tasks)
	if err.Error != nil {
		return nil, err.Error
	}
	return tasks, nil
}

func FirstTask(db *gorm.DB, id string) (Task, error) {
	var task Task
	err := db.First(&task, id)
	if err.Error != nil {
		return task, err.Error
	}
	return task, nil
}

func CreateTask(db *gorm.DB, task Task) (Task, error) {
	err := db.Create(&task)
	if err.Error != nil {
		return task, err.Error
	}
	return task, nil
}

func DeleteTask(db *gorm.DB, id string) (Task, error) {
	var task Task
	err := db.Delete(&task, id)
	if err.Error != nil {
		return task, err.Error
	}
	return task, nil
}
