package main

// go test -v main_test.go

import (
	"testing"

	"goapi/database"
	"goapi/model"
)

func TestGetAllTasks(t *testing.T) {
	database.InitDb()
	var tasks []model.Task
	err := database.Db.Find(&tasks).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetTask(t *testing.T) {
	database.InitDb()
	var task model.Task
	err := database.Db.First(&task).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestAddTask(t *testing.T) {
	database.InitDb()
	task := &model.Task{
		Text: "Add Text for Test",
		Done: false,
	}
	err := database.Db.Create(&task).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdateTask(t *testing.T) {
	database.InitDb()
	var task model.Task
	database.Db.First(&task, "0")
	task.Text = "Update Text for Test"
	err := database.Db.Save(&task).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeleteTask(t *testing.T) {
	database.InitDb()
	var task model.Task
	err := database.Db.Delete(&task, "0").Error
	if err != nil {
		t.Fatal(err.Error())
	}
}
