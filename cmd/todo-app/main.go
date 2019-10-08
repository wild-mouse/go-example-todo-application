package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
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
func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/tasks/", handlers.MakeHandler(handlers.TasksHandler, db))

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"),
	))
	http.ListenAndServe(":1323", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
