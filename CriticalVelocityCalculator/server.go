package main

import (
	"fmt"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("/Users/lky/code_repos/go/src/goStudy/CriticalVelocityCalculator/*")

	router.GET("/",
		func(c *gin.Context) {
			c.HTML(200, "CVC.html", gin.H{})
		})

	err := router.Run(":80")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if nil != err {
		fmt.Errorf("%s", err)
		panic(err)
	}
}
