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

// GetTask implements get_task.
func (s *taskssrvc) GetTask(ctx context.Context, p *tasks.GetTaskPayload) (res *tasks.Task, err error) {
	res = &tasks.Task{}
	s.logger.Print("tasks.get_task")
	return
}

// GetTasks implements get_tasks.
func (s *taskssrvc) GetTasks(ctx context.Context) (res []*tasks.Task, err error) {
	s.logger.Print("tasks.get_tasks")
	return
}

// AddTask implements add_task.
func (s *taskssrvc) AddTask(ctx context.Context, p *tasks.Task) (res *tasks.Task, err error) {
	res = &tasks.Task{}
	s.logger.Print("tasks.add_task")
	return
}

// UpdateTask implements update_task.
func (s *taskssrvc) UpdateTask(ctx context.Context, p *tasks.Task) (res *tasks.Task, err error) {
	res = &tasks.Task{}
	s.logger.Print("tasks.update_task")
	return
}

// DeleteTask implements delete_task.
func (s *taskssrvc) DeleteTask(ctx context.Context, p *tasks.DeleteTaskPayload) (res *tasks.Task, err error) {
	res = &tasks.Task{}
	s.logger.Print("tasks.delete_task")
	return
}
