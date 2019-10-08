package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/wild-mouse/go-example-todo-application/pkg/models"
	"log"
	"net/http"
)

// GetTasks godoc
// @Summary Show a tasks
// @Description get all tasks
// @ID get-string-by-int
// @Produce  json
// @Success 200 {string} string
// @Router /tasks/ [get]
func GetTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	query := fmt.Sprintf("SELECT * FROM tasks")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
	defer rows.Close()
	tasks := make([]models.Task, 0)

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.Id, &task.Name); err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	bytes, _ := json.Marshal(tasks)
	_, err = fmt.Fprint(w, string(bytes))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetTask godoc
// @Summary Get a task
// @Description get task by ID
// @ID get-task-by-int
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Router /tasks/{id} [get]
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
	if err != nil {
		fmt.Println(err)
		return
	}
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

func UpdateTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	decoder := json.NewDecoder(r.Body)
	var updateTask models.Task
	err := decoder.Decode(&updateTask)
	if err != nil {
		fmt.Println(err)
		return
	}
	if updateTask.Name == "" {
		http.Error(w, "Task name shouldn't be empty.", http.StatusBadRequest)
		return
	}
	id := r.URL.Path[len("/tasks/"):]
	query := fmt.Sprintf("UPDATE tasks SET name = \"%s\" WHERE id = \"%s\"", updateTask.Name, id)
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
		http.Error(w, "Tasks to be updated not found.", http.StatusNotFound)
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
