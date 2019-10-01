package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tasks", func() {
	Title("TODO Task Service")
	Description("Service for tasks")
	Server("tasks", func() {
		Services("tasks")
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
			Response(CodeOK)
		})
	})

	Method("get_tasks", func() {
		Result(ArrayOf(Task))
		HTTP(func() {
			GET("/tasks")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add_task", func() {
		Payload(Task)
		Result(Task)
		HTTP(func() {
			POST("/tasks")
			Response(StatusCreated)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("update_task", func() {
		Payload(Task)
		Result(Task)
		HTTP(func() {
			PUT("/tasks/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("delete_task", func() {
		Payload(func() {
			Field(1, "id", String)
			Required("id")
		})
		Result(Task)
		HTTP(func() {
			DELETE("/task/{id}")
			Response(StatusOK)
		})
		GRPC(func() {
			Response(CodeOK)
		})
	})
})
