// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC client types
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package client

import (
	"unicode/utf8"

	taskspb "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/pb"
	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
	goa "goa.design/goa/v3/pkg"
)

// NewGetTaskRequest builds the gRPC request type from the payload of the
// "get_task" endpoint of the "tasks" service.
func NewGetTaskRequest(payload *tasks.GetTaskPayload) *taskspb.GetTaskRequest {
	message := &taskspb.GetTaskRequest{
		Id: payload.ID,
	}
	return message
}

// NewGetTaskResult builds the result type of the "get_task" endpoint of the
// "tasks" service from the gRPC response type.
func NewGetTaskResult(message *taskspb.GetTaskResponse) *tasks.Task {
	result := &tasks.Task{
		Name: message.Name,
	}
	if message.Id != 0 {
		result.ID = &message.Id
	}
	return result
}

// NewGetTasksRequest builds the gRPC request type from the payload of the
// "get_tasks" endpoint of the "tasks" service.
func NewGetTasksRequest() *taskspb.GetTasksRequest {
	message := &taskspb.GetTasksRequest{}
	return message
}

// NewGetTasksResult builds the result type of the "get_tasks" endpoint of the
// "tasks" service from the gRPC response type.
func NewGetTasksResult(message *taskspb.GetTasksResponse) []*tasks.Task {
	result := make([]*tasks.Task, len(message.Field))
	for i, val := range message.Field {
		result[i] = &tasks.Task{
			Name: val.Name,
		}
		if val.Id != 0 {
			result[i].ID = &val.Id
		}
	}
	return result
}

// NewAddTaskRequest builds the gRPC request type from the payload of the
// "add_task" endpoint of the "tasks" service.
func NewAddTaskRequest(payload *tasks.Task) *taskspb.AddTaskRequest {
	message := &taskspb.AddTaskRequest{
		Name: payload.Name,
	}
	if payload.ID != nil {
		message.Id = *payload.ID
	}
	return message
}

// NewAddTaskResult builds the result type of the "add_task" endpoint of the
// "tasks" service from the gRPC response type.
func NewAddTaskResult(message *taskspb.AddTaskResponse) *tasks.Task {
	result := &tasks.Task{
		Name: message.Name,
	}
	if message.Id != 0 {
		result.ID = &message.Id
	}
	return result
}

// NewUpdateTaskRequest builds the gRPC request type from the payload of the
// "update_task" endpoint of the "tasks" service.
func NewUpdateTaskRequest(payload *tasks.Task) *taskspb.UpdateTaskRequest {
	message := &taskspb.UpdateTaskRequest{
		Name: payload.Name,
	}
	if payload.ID != nil {
		message.Id = *payload.ID
	}
	return message
}

// NewUpdateTaskResult builds the result type of the "update_task" endpoint of
// the "tasks" service from the gRPC response type.
func NewUpdateTaskResult(message *taskspb.UpdateTaskResponse) *tasks.Task {
	result := &tasks.Task{
		Name: message.Name,
	}
	if message.Id != 0 {
		result.ID = &message.Id
	}
	return result
}

// NewDeleteTaskRequest builds the gRPC request type from the payload of the
// "delete_task" endpoint of the "tasks" service.
func NewDeleteTaskRequest(payload *tasks.DeleteTaskPayload) *taskspb.DeleteTaskRequest {
	message := &taskspb.DeleteTaskRequest{
		Id: payload.ID,
	}
	return message
}

// NewDeleteTaskResult builds the result type of the "delete_task" endpoint of
// the "tasks" service from the gRPC response type.
func NewDeleteTaskResult(message *taskspb.DeleteTaskResponse) *tasks.Task {
	result := &tasks.Task{
		Name: message.Name,
	}
	if message.Id != 0 {
		result.ID = &message.Id
	}
	return result
}

// ValidateGetTaskResponse runs the validations defined on GetTaskResponse.
func ValidateGetTaskResponse(message *taskspb.GetTaskResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	return
}

// ValidateGetTasksResponse runs the validations defined on GetTasksResponse.
func ValidateGetTasksResponse(message *taskspb.GetTasksResponse) (err error) {
	if message.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("field", "message"))
	}
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateTask(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateTask runs the validations defined on Task.
func ValidateTask(message *taskspb.Task) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	return
}

// ValidateAddTaskResponse runs the validations defined on AddTaskResponse.
func ValidateAddTaskResponse(message *taskspb.AddTaskResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	return
}

// ValidateUpdateTaskResponse runs the validations defined on
// UpdateTaskResponse.
func ValidateUpdateTaskResponse(message *taskspb.UpdateTaskResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	return
}

// ValidateDeleteTaskResponse runs the validations defined on
// DeleteTaskResponse.
func ValidateDeleteTaskResponse(message *taskspb.DeleteTaskResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	return
}
