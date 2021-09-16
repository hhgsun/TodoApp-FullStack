package main

// go test -v main_test.go

import (
	"fmt"

	"github.com/hhgsun/TodoApp-FullStack/api/database"
	"github.com/hhgsun/TodoApp-FullStack/api/routers"
)

func main() {
	fmt.Println("Hi!")
	database.InitDb()
	routers.HandleRequests()
}
