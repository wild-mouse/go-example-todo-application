// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks HTTP server types
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package server

import (
	"unicode/utf8"

	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
	goa "goa.design/goa/v3/pkg"
)

// AddTaskRequestBody is the type of the "tasks" service "add_task" endpoint
// HTTP request body.
type AddTaskRequestBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// UpdateTaskRequestBody is the type of the "tasks" service "update_task"
// endpoint HTTP request body.
type UpdateTaskRequestBody struct {
	// Name of task
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// GetTaskResponseBody is the type of the "tasks" service "get_task" endpoint
// HTTP response body.
type GetTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// GetTasksResponseBody is the type of the "tasks" service "get_tasks" endpoint
// HTTP response body.
type GetTasksResponseBody []*TaskResponse

// AddTaskResponseBody is the type of the "tasks" service "add_task" endpoint
// HTTP response body.
type AddTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// UpdateTaskResponseBody is the type of the "tasks" service "update_task"
// endpoint HTTP response body.
type UpdateTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// DeleteTaskResponseBody is the type of the "tasks" service "delete_task"
// endpoint HTTP response body.
type DeleteTaskResponseBody struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// TaskResponse is used to define fields on response body types.
type TaskResponse struct {
	// ID is the unique id of the task.
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of task
	Name string `form:"name" json:"name" xml:"name"`
}

// NewGetTaskResponseBody builds the HTTP response body from the result of the
// "get_task" endpoint of the "tasks" service.
func NewGetTaskResponseBody(res *tasks.Task) *GetTaskResponseBody {
	body := &GetTaskResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewGetTasksResponseBody builds the HTTP response body from the result of the
// "get_tasks" endpoint of the "tasks" service.
func NewGetTasksResponseBody(res []*tasks.Task) GetTasksResponseBody {
	body := make([]*TaskResponse, len(res))
	for i, val := range res {
		body[i] = &TaskResponse{
			ID:   val.ID,
			Name: val.Name,
		}
	}
	return body
}

// NewAddTaskResponseBody builds the HTTP response body from the result of the
// "add_task" endpoint of the "tasks" service.
func NewAddTaskResponseBody(res *tasks.Task) *AddTaskResponseBody {
	body := &AddTaskResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewUpdateTaskResponseBody builds the HTTP response body from the result of
// the "update_task" endpoint of the "tasks" service.
func NewUpdateTaskResponseBody(res *tasks.Task) *UpdateTaskResponseBody {
	body := &UpdateTaskResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewDeleteTaskResponseBody builds the HTTP response body from the result of
// the "delete_task" endpoint of the "tasks" service.
func NewDeleteTaskResponseBody(res *tasks.Task) *DeleteTaskResponseBody {
	body := &DeleteTaskResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewGetTaskPayload builds a tasks service get_task endpoint payload.
func NewGetTaskPayload(id uint32) *tasks.GetTaskPayload {
	return &tasks.GetTaskPayload{
		ID: id,
	}
}

// NewAddTaskTask builds a tasks service add_task endpoint payload.
func NewAddTaskTask(body *AddTaskRequestBody) *tasks.Task {
	v := &tasks.Task{
		ID:   body.ID,
		Name: *body.Name,
	}
	return v
}

// NewUpdateTaskTask builds a tasks service update_task endpoint payload.
func NewUpdateTaskTask(body *UpdateTaskRequestBody, id uint32) *tasks.Task {
	v := &tasks.Task{
		Name: *body.Name,
	}
	v.ID = &id
	return v
}

// NewDeleteTaskPayload builds a tasks service delete_task endpoint payload.
func NewDeleteTaskPayload(id string) *tasks.DeleteTaskPayload {
	return &tasks.DeleteTaskPayload{
		ID: id,
	}
}

// ValidateAddTaskRequestBody runs the validations defined on
// add_task_request_body
func ValidateAddTaskRequestBody(body *AddTaskRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateUpdateTaskRequestBody runs the validations defined on
// update_task_request_body
func ValidateUpdateTaskRequestBody(body *UpdateTaskRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	return
}

// ValidateTaskResponse runs the validations defined on taskResponse
func ValidateTaskResponse(body *TaskResponse) (err error) {
	if utf8.RuneCountInString(body.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
	}
	return
}
