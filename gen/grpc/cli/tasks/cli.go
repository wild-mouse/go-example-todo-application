// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks gRPC client CLI support package
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package cli

import (
	"flag"
	"fmt"
	"os"

	tasksc "github.com/wild-mouse/go-example-todo-application/gen/grpc/tasks/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `tasks (get-task|get-tasks|add-task|update-task|delete-task)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` tasks get-task --message '{
      "id": 955004123
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		tasksFlags = flag.NewFlagSet("tasks", flag.ContinueOnError)

		tasksGetTaskFlags       = flag.NewFlagSet("get-task", flag.ExitOnError)
		tasksGetTaskMessageFlag = tasksGetTaskFlags.String("message", "", "")

		tasksGetTasksFlags = flag.NewFlagSet("get-tasks", flag.ExitOnError)

		tasksAddTaskFlags       = flag.NewFlagSet("add-task", flag.ExitOnError)
		tasksAddTaskMessageFlag = tasksAddTaskFlags.String("message", "", "")

		tasksUpdateTaskFlags       = flag.NewFlagSet("update-task", flag.ExitOnError)
		tasksUpdateTaskMessageFlag = tasksUpdateTaskFlags.String("message", "", "")

		tasksDeleteTaskFlags       = flag.NewFlagSet("delete-task", flag.ExitOnError)
		tasksDeleteTaskMessageFlag = tasksDeleteTaskFlags.String("message", "", "")
	)
	tasksFlags.Usage = tasksUsage
	tasksGetTaskFlags.Usage = tasksGetTaskUsage
	tasksGetTasksFlags.Usage = tasksGetTasksUsage
	tasksAddTaskFlags.Usage = tasksAddTaskUsage
	tasksUpdateTaskFlags.Usage = tasksUpdateTaskUsage
	tasksDeleteTaskFlags.Usage = tasksDeleteTaskUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "tasks":
			svcf = tasksFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "tasks":
			switch epn {
			case "get-task":
				epf = tasksGetTaskFlags

			case "get-tasks":
				epf = tasksGetTasksFlags

			case "add-task":
				epf = tasksAddTaskFlags

			case "update-task":
				epf = tasksUpdateTaskFlags

			case "delete-task":
				epf = tasksDeleteTaskFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "tasks":
			c := tasksc.NewClient(cc, opts...)
			switch epn {
			case "get-task":
				endpoint = c.GetTask()
				data, err = tasksc.BuildGetTaskPayload(*tasksGetTaskMessageFlag)
			case "get-tasks":
				endpoint = c.GetTasks()
				data = nil
			case "add-task":
				endpoint = c.AddTask()
				data, err = tasksc.BuildAddTaskPayload(*tasksAddTaskMessageFlag)
			case "update-task":
				endpoint = c.UpdateTask()
				data, err = tasksc.BuildUpdateTaskPayload(*tasksUpdateTaskMessageFlag)
			case "delete-task":
				endpoint = c.DeleteTask()
				data, err = tasksc.BuildDeleteTaskPayload(*tasksDeleteTaskMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// tasksUsage displays the usage of the tasks command and its subcommands.
func tasksUsage() {
	fmt.Fprintf(os.Stderr, `The task service.
Usage:
    %s [globalflags] tasks COMMAND [flags]

COMMAND:
    get-task: GetTask implements get_task.
    get-tasks: GetTasks implements get_tasks.
    add-task: AddTask implements add_task.
    update-task: UpdateTask implements update_task.
    delete-task: DeleteTask implements delete_task.

Additional help:
    %s tasks COMMAND --help
`, os.Args[0], os.Args[0])
}
func tasksGetTaskUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks get-task -message JSON

GetTask implements get_task.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks get-task --message '{
      "id": 955004123
   }'
`, os.Args[0])
}

func tasksGetTasksUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks get-tasks

GetTasks implements get_tasks.

Example:
    `+os.Args[0]+` tasks get-tasks
`, os.Args[0])
}

func tasksAddTaskUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks add-task -message JSON

AddTask implements add_task.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks add-task --message '{
      "id": 1,
      "name": "Implement awesome application using Goa"
   }'
`, os.Args[0])
}

func tasksUpdateTaskUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks update-task -message JSON

UpdateTask implements update_task.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks update-task --message '{
      "id": 1,
      "name": "Implement awesome application using Goa"
   }'
`, os.Args[0])
}

func tasksDeleteTaskUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] tasks delete-task -message JSON

DeleteTask implements delete_task.
    -message JSON: 

Example:
    `+os.Args[0]+` tasks delete-task --message '{
      "id": "Quam voluptatum."
   }'
`, os.Args[0])
}
