package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

var err error

func Connection() {
	connstr := os.Getenv("DB_CONN")
	if connstr == "" {
		connstr = "postgres://postgres:hhgsun@localhost:5432/todo_db?sslmode=disable"
	}
	fmt.Println("CONNECT STR: " + connstr)
	db, err = gorm.Open("postgres", connstr) //"postgres://postgres:hhgsun@localhost:5432/todo_db?sslmode=disable"
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// defer db.Close()

	db.AutoMigrate(&Task{})
}
