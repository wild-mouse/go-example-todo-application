package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wild-mouse/go-example-todo-application/pkg/models"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	row := db.QueryRow("SELECT * FROM tasks WHERE id=1")
	var task models.Task
	err := row.Scan(&task.Id, &task.Name)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(task)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		fmt.Println(err)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("Handling POST Method")
	_, err := db.Query("insert into tasks values (1, \"hogehoge\")")
	if err != nil {
		fmt.Println(err)
	}
}

