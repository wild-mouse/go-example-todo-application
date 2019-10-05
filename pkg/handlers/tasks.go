package handlers

import (
	"database/sql"
	"fmt"
	"github.com/wild-mouse/go-example-todo-application/pkg/services"
	"net/http"
)

func MakeHandler(fn func(http.ResponseWriter, *http.Request, *sql.DB), db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, db)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		if r.URL.Path == "/tasks/" {
			fmt.Println("Getting All tasks")
		} else {
			services.GetTask(w, r, db)
		}
	}
	if r.Method == http.MethodPost {
		services.AddTask(w, r, db)
	}
	if r.Method == http.MethodPut {
		fmt.Println("Handling PUT Method")
	}
	if r.Method == http.MethodDelete {
		services.DeleteTask(w, r, db)
	}
}

