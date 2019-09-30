package tasksapi

import (
	"context"
	"log"

	"github.com/wild-mouse/go-example-todo-application/gen/tasks"
)

// tasks service example implementation.
// The example methods log the requests and return zero values.
type taskssrvc struct {
	logger *log.Logger
}

// NewTasks returns the tasks service implementation.
func NewTasks(logger *log.Logger) tasks.Service {
	return &taskssrvc{logger}
}

// GetTask implements get_task.
func (s *taskssrvc) GetTask(ctx context.Context, p *tasks.GetTaskPayload) (res *tasks.Task, err error) {
	id := uint32(1);
	res = &tasks.Task{ID: &id, Name: "This is a simple task."}
	s.logger.Print("tasks.get_task")
	return res, nil
}
