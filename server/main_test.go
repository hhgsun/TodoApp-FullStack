package main

import (
	"testing"
)

// go test -v main_test.go main.go

func TestAddTask(t *testing.T) {
	task := Task{
		Text: "adasd",
	}
	if task.ID > 0 {
		t.Error("Task Eklemede Hata")
	}
}
