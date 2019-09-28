package main

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func main() {
	task := &Task{100, "Task name"}
	bytes, _ := json.Marshal(task)
	fmt.Println(string(bytes))
}
