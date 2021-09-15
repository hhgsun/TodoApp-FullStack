package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Text string `json:"Text"`
	Done bool   `json:"Done"`
}

type Tasks []Task
