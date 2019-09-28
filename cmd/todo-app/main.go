package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/wild-mouse/go-example-todo-application/pkg/models"
)

func makeHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		getTask(w, r, db)
	}
	if r.Method == http.MethodPost {
		addTask(w, r, db)
	}
	if r.Method == http.MethodPut {
		fmt.Println("Handling PUT Method")
	}
	if r.Method == http.MethodDelete {
		fmt.Println("Handling DELETE Method")
	}
}

func getTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

func addTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Println("Handling POST Method")
	_, err := db.Query("insert into tasks values (1, \"hogehoge\")")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/tasks", makeHandler(tasksHandler, db))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
