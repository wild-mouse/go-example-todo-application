// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC server encoders and decoders
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package server

import (
	"context"

	taskspb "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/pb"
	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeGetTaskResponse encodes responses from the "tasks" service "get_task"
// endpoint.
func EncodeGetTaskResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*tasks.Task)
	if !ok {
		return nil, goagrpc.ErrInvalidType("tasks", "get_task", "*tasks.Task", v)
	}
	resp := NewGetTaskResponse(result)
	return resp, nil
}

// DecodeGetTaskRequest decodes requests sent to "tasks" service "get_task"
// endpoint.
func DecodeGetTaskRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *taskspb.GetTaskRequest
		ok      bool
	)
	{
		if message, ok = v.(*taskspb.GetTaskRequest); !ok {
			return nil, goagrpc.ErrInvalidType("tasks", "get_task", "*taskspb.GetTaskRequest", v)
		}
	}
	var payload *tasks.GetTaskPayload
	{
		payload = NewGetTaskPayload(message)
	}
	return payload, nil
}
