package main

import (
	//"github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"net/http"
)

func main() {
	db, err := sql.Open("mysql", "docker:docker@/playground")
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Query("insert into tasks values (1, \"hogehoge\")")
	if err != nil {
		fmt.Println(err)
	}

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()
}
