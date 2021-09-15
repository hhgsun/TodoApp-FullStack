package main

// go test -v main_test.go

import (
	"fmt"

	"github.com/hhgsun/goapi/database"
	"github.com/hhgsun/goapi/routers"
)

func main() {
	fmt.Println("Hi!")
	database.InitDb()
	routers.HandleRequests()
}
