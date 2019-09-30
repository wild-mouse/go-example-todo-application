package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tasks", func() {
	Title("TODO Task Service")
	Description("Service for tasks")
	Server("tasks", func() {
		Services("tasks", "swagger")
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var Task = Type("task", func() {
	Description("Task describes a task to be acted.")
	Attribute("id", UInt32, "ID is the unique id of the task.", func() {
		Example(1)
		Meta("rpc:tag", "1")
	})
	Attribute("name", String, "Name of task", func() {
		MaxLength(100)
		Example("Implement awesome application using Goa")
		Meta("rpc:tag", "2")
	})
	Required("name")
})

var _ = Service("tasks", func() {
	Description("The task service.")

	Method("get_task", func() {
		Payload(func() {
			Field(1, "id", UInt32, "ID of task")
			Required("id")
		})
		Result(Task)
		HTTP(func() {
			GET("/tasks/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})
	Files("/swagger.json", "gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})
