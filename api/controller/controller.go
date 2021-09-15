package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hhgsun/goapi/database"
	"github.com/hhgsun/goapi/model"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []model.Task
	err := database.Db.Find(&tasks)
	if err.Error != nil {
		json.NewEncoder(w).Encode(err.Error.Error())
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task model.Task
	err := database.Db.First(&task, params["id"])
	if err.Error != nil {
		json.NewEncoder(w).Encode(err.Error.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)
	err := database.Db.Create(&task)
	if err.Error != nil {
		json.NewEncoder(w).Encode(err.Error.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task model.Task
	database.Db.First(&task, params["id"])
	json.NewDecoder(r.Body).Decode(&task)
	err := database.Db.Save(&task)
	if err.Error != nil {
		json.NewEncoder(w).Encode(err.Error.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task model.Task
	err := database.Db.Delete(&task, params["id"])
	if err.Error != nil {
		json.NewEncoder(w).Encode(err.Error.Error())
	} else {
		json.NewEncoder(w).Encode(true)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("To Do Api: /tasks")
}
