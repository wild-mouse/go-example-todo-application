package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wild-mouse/go-example-todo-application/pkg/models"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.URL.Path[len("/tasks/"):]
	query := fmt.Sprintf("SELECT * FROM tasks WHERE id=%s", id)
	fmt.Println(query)
	row := db.QueryRow(query)
	var task models.Task
	err := row.Scan(&task.Id, &task.Name)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Task not found.", http.StatusNotFound)
		return
	}
	bytes, _ := json.Marshal(task)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("Handling POST Method")
	_, err := db.Query("insert into tasks values (1, \"hogehoge\")")
	if err != nil {
		fmt.Println(err)
	}
}

