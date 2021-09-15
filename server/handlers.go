package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Text string `json:"Text"`
	Done bool   `json:"Done"`
}

type Tasks []Task

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task
	db.Find(&tasks)
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	db.Create(&task)
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&task)
	db.Save(&task)
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task Task
	db.Delete(&task, params["id"])
	if err == gorm.ErrRecordNotFound {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(true)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("To Do Api: /tasks")
}
