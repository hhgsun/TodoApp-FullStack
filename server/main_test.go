package main

// go test -v main_test.go

import (
	"testing"
)

func TestAllTasks(t *testing.T) {
	if 2 > 3 {
		t.Error("Büyük Dikkat")
	}
}

/* func AddTaskTest(t *testing.T, db *gorm.DB) {
	task := Task{Text: "random", Done: false}
	if err := db.Create(&task).Error; err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		db.Delete(&task)
	})
}

func TestConnection(t *testing.T) {
	connstr := os.Getenv("DB_CONN")
	if connstr == "" {
		connstr = "postgres://postgres:hhgsun@localhost:5432/todo_db?sslmode=disable"
	}
	db, err = gorm.Open("postgres", connstr)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	db.AutoMigrate(&Task{})
}
*/
