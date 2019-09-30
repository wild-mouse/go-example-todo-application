// Code generated by goa v3.0.6, DO NOT EDIT.
//
// HTTP request path constructors for the tasks service.
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package server

import (
	"fmt"
)

// GetTaskTasksPath returns the URL path to the tasks service get_task HTTP endpoint.
func GetTaskTasksPath(id uint32) string {
	return fmt.Sprintf("/tasks/%v", id)
}
