// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks client HTTP transport
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the tasks service endpoint HTTP clients.
type Client struct {
	// GetTask Doer is the HTTP client used to make requests to the get_task
	// endpoint.
	GetTaskDoer goahttp.Doer

	// GetTasks Doer is the HTTP client used to make requests to the get_tasks
	// endpoint.
	GetTasksDoer goahttp.Doer

	// AddTask Doer is the HTTP client used to make requests to the add_task
	// endpoint.
	AddTaskDoer goahttp.Doer

	// UpdateTask Doer is the HTTP client used to make requests to the update_task
	// endpoint.
	UpdateTaskDoer goahttp.Doer

	// DeleteTask Doer is the HTTP client used to make requests to the delete_task
	// endpoint.
	DeleteTaskDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the tasks service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetTaskDoer:         doer,
		GetTasksDoer:        doer,
		AddTaskDoer:         doer,
		UpdateTaskDoer:      doer,
		DeleteTaskDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetTask returns an endpoint that makes HTTP requests to the tasks service
// get_task server.
func (c *Client) GetTask() goa.Endpoint {
	var (
		decodeResponse = DecodeGetTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTaskDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("tasks", "get_task", err)
		}
		return decodeResponse(resp)
	}
}

// GetTasks returns an endpoint that makes HTTP requests to the tasks service
// get_tasks server.
func (c *Client) GetTasks() goa.Endpoint {
	var (
		decodeResponse = DecodeGetTasksResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetTasksRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetTasksDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("tasks", "get_tasks", err)
		}
		return decodeResponse(resp)
	}
}

// AddTask returns an endpoint that makes HTTP requests to the tasks service
// add_task server.
func (c *Client) AddTask() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddTaskRequest(c.encoder)
		decodeResponse = DecodeAddTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddTaskDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("tasks", "add_task", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateTask returns an endpoint that makes HTTP requests to the tasks service
// update_task server.
func (c *Client) UpdateTask() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateTaskRequest(c.encoder)
		decodeResponse = DecodeUpdateTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateTaskDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("tasks", "update_task", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteTask returns an endpoint that makes HTTP requests to the tasks service
// delete_task server.
func (c *Client) DeleteTask() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteTaskResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteTaskRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteTaskDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("tasks", "delete_task", err)
		}
		return decodeResponse(resp)
	}
}
