package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/wild-mouse/go-example-todo-application/docs"
	"github.com/wild-mouse/go-example-todo-application/pkg/handlers"
	"log"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService
// @host localhost:8080
func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/tasks/", handlers.MakeHandler(handlers.TasksHandler, db))
	http.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
