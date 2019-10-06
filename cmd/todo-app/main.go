package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wild-mouse/go-example-todo-application/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/tasks/", handlers.MakeHandler(handlers.TasksHandler, db))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
