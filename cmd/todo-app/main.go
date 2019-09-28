package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"github.com/wild-mouse/go-example-todo-application/pkg/handlers"
)

func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/tasks", handlers.MakeHandler(handlers.TasksHandler, db))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
