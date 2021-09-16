package controller

import (
	"encoding/json"
	"net/http"

	"goapi/database"
	"goapi/model"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err2 := model.FindTasks(database.Db)
	if err2 != nil {
		json.NewEncoder(w).Encode(err2.Error())
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	task, err := model.FirstTask(database.Db, params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(task)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task model.Task
	json.NewDecoder(r.Body).Decode(&task)
	task, err := model.CreateTask(database.Db, task)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
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
	_, err := model.DeleteTask(database.Db, params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(true)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("To Do Api: /tasks")
}
