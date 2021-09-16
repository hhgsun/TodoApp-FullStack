package main

// go test -v main_test.go

import (
	"testing"

	"goapi/database"
	"goapi/model"

	"github.com/jinzhu/gorm"
)

func TestGetAllTasks(t *testing.T) {
	database.InitDb()
	_, err := model.FindTasks(database.Db)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetTask(t *testing.T) {
	database.InitDb()
	_, err := model.FirstTask(database.Db, "0")
	if err != gorm.ErrRecordNotFound {
		t.Fatal(err.Error())
	}
}

func TestAddTask(t *testing.T) {
	database.InitDb()
	task := &model.Task{
		Text: "Add Text for Test",
		Done: false,
	}
	_, err := model.CreateTask(database.Db, *task)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdateTask(t *testing.T) {
	database.InitDb()
	var task model.Task
	database.Db.First(&task, "0")
	task.Text = "Update Text for Test"
	err := database.Db.Save(&task).Error
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeleteTask(t *testing.T) {
	database.InitDb()
	_, err := model.DeleteTask(database.Db, "0")
	if err != nil {
		t.Fatal(err.Error())
	}
}

/* import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	_ "github.com/lib/pq"
)

func insertTask(db *sql.DB, ID int64, Text string, Done bool) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("INSERT INTO task_viewers (ID, Text, Done) VALUES (?, ?, ?)", ID, Text, Done, false); err != nil {
		return
	}
	return
}

// a successful case
func TestUpdateDoneAndText(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	columns := []string{"ID", "Text", "Done"}
	// expect transaction begin
	mock.ExpectBegin()
	// expect query to fetch order and user, match it with regexp
	mock.ExpectQuery("SELECT (.+) FROM tasks").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Test Text", false))
	// expect user balance update
	mock.ExpectPrepare("UPDATE users SET Done").ExpectExec().
		WithArgs(true, 1).                        // refund amount, user id
		WillReturnResult(sqlmock.NewResult(0, 1)) // no insert id, 1 affected row
	// expect order status update
	mock.ExpectPrepare("UPDATE orders SET Text").ExpectExec().
		WithArgs("Text", 1).                      // status, id
		WillReturnResult(sqlmock.NewResult(0, 1)) // no insert id, 1 affected row
	// expect a transaction commit
	mock.ExpectCommit()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateDoneAndText(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE tasks").WillReturnResult(sqlmock.NewResult(1, 1))
	//mock.ExpectExec("INSERT INTO task_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = insertTask(db, 0, "Example", false); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
} */
