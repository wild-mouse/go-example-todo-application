package handlers

import (
	"database/sql"
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
			services.GetTasks(w, r, db)
		} else {
			services.GetTask(w, r, db)
		}
	}
	if r.Method == http.MethodPost {
		if r.URL.Path == "/tasks/" {
			services.AddTask(w, r, db)
		} else {
			http.NotFound(w, r)
		}
	}
	if r.Method == http.MethodPut {
		if r.URL.Path == "/tasks/" {
			http.NotFound(w, r)
		} else {
			services.UpdateTask(w, r, db)
		}
	}
	if r.Method == http.MethodDelete {
		if r.URL.Path == "/tasks/" {
			http.NotFound(w, r)
		} else {
			services.DeleteTask(w, r, db)
		}
	}
}

