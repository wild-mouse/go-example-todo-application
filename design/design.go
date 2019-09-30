package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("tasks", func() {
	Title("TODO Task Service")
	Description("Service for tasks")
	Server("tasks", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("tasks", func() {
	Description("The task service.")

	Method("count_tasks", func() {
		Result(Int)
		HTTP(func() {
			GET("/tasks/count")
			Response(StatusOK)
		})
		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
