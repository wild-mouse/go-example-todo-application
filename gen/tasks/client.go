// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks client
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package tasks

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "tasks" service client.
type Client struct {
	GetTaskEndpoint    goa.Endpoint
	GetTasksEndpoint   goa.Endpoint
	AddTaskEndpoint    goa.Endpoint
	UpdateTaskEndpoint goa.Endpoint
	DeleteTaskEndpoint goa.Endpoint
}

// NewClient initializes a "tasks" service client given the endpoints.
func NewClient(getTask, getTasks, addTask, updateTask, deleteTask goa.Endpoint) *Client {
	return &Client{
		GetTaskEndpoint:    getTask,
		GetTasksEndpoint:   getTasks,
		AddTaskEndpoint:    addTask,
		UpdateTaskEndpoint: updateTask,
		DeleteTaskEndpoint: deleteTask,
	}
}

// GetTask calls the "get_task" endpoint of the "tasks" service.
func (c *Client) GetTask(ctx context.Context, p *GetTaskPayload) (res *Task, err error) {
	var ires interface{}
	ires, err = c.GetTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Task), nil
}

// GetTasks calls the "get_tasks" endpoint of the "tasks" service.
func (c *Client) GetTasks(ctx context.Context) (res []*Task, err error) {
	var ires interface{}
	ires, err = c.GetTasksEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Task), nil
}

// AddTask calls the "add_task" endpoint of the "tasks" service.
func (c *Client) AddTask(ctx context.Context, p *Task) (res *Task, err error) {
	var ires interface{}
	ires, err = c.AddTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Task), nil
}

// UpdateTask calls the "update_task" endpoint of the "tasks" service.
func (c *Client) UpdateTask(ctx context.Context, p *Task) (res *Task, err error) {
	var ires interface{}
	ires, err = c.UpdateTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Task), nil
}

// DeleteTask calls the "delete_task" endpoint of the "tasks" service.
func (c *Client) DeleteTask(ctx context.Context, p *DeleteTaskPayload) (res *Task, err error) {
	var ires interface{}
	ires, err = c.DeleteTaskEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Task), nil
}
