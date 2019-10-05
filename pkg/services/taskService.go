package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wild-mouse/go-example-todo-application/pkg/models"
	"net/http"
)

func GetTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("Getting Tasks.")
}

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
	decoder := json.NewDecoder(r.Body)
	var newTask models.Task
	err := decoder.Decode(&newTask)
	if newTask.Name == "" {
		http.Error(w, "Task name shouldn't be empty.", http.StatusBadRequest)
		return
	}
	query := fmt.Sprintf("insert into tasks (name) values (\"%s\")", newTask.Name)
	_, err = db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteTask( w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := r.URL.Path[len("/tasks/"):]
	query := fmt.Sprintf("DELETE FROM tasks WHERE id=%s", id)
	result, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	if count == 0 {
		http.Error(w, "Tasks to be deleted not found.", http.StatusNotFound)
		return
	}
}
