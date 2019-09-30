package tasksapi

import (
	"context"
	"log"

	tasks "github.com/wild-mouse/go-example-todo-application/gen/tasks"
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

// CountTasks implements count_tasks.
func (s *taskssrvc) CountTasks(ctx context.Context) (res int, err error) {
	s.logger.Print("tasks.count_tasks")
	return
}
