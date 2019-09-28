package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Task struct {
	Id uint32
	Name string
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Println("Handling GET Method")
	}
	if r.Method == http.MethodPost {
		fmt.Println("Handling POST Method")
	}
	if r.Method == http.MethodPut {
		fmt.Println("Handling PUT Method")
	}
	if r.Method == http.MethodDelete {
		fmt.Println("Handling DELETE Method")
	}
}

func main() {
	//db, err := sql.Open("mysql", "docker:docker@/playground")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//_, err = db.Query("insert into tasks values (1, \"hogehoge\")")
	//if err != nil {
	//	fmt.Println(err)
	//}
	http.HandleFunc("/tasks", tasksHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()
}
