package main

// go test -v main_test.go

import (
	"fmt"

	"goapi/database"
	"goapi/routers"
)

func main() {
	fmt.Println("Hi!")

	database.InitDb()
	routers.HandleRequests()
}
