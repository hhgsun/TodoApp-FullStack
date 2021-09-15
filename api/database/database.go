package database

import (
	"fmt"
	"os"

	"github.com/hhgsun/goapi/model"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var (
	Db  *gorm.DB
	Err error
)

func InitDb() {
	connstr := os.Getenv("DB_CONN")
	if connstr == "" {
		connstr = "postgres://postgres:hhgsun@localhost:5432/todo_db?sslmode=disable"
	}
	//fmt.Println("CONNECT STR: " + connstr)
	Db, Err = gorm.Open("postgres", connstr) //"postgres://postgres:hhgsun@localhost:5432/todo_db?sslmode=disable"
	if Err != nil {
		fmt.Println(Err.Error())
		panic("failed to connect database")
	}

	// defer db.Close()

	Db.AutoMigrate(&model.Task{})
}
