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
			keys := []string{
				"Distance0","Distance1","Distance2",
				"RaceTime0","RaceTime1","RaceTime2",
			}

			for _, key := range keys {
				fmt.Println(c.Query(key))
			}
		})

	// router.POST("/",
	// 	func(c *gin.Context) {
	// 		c.HTML(200, "CVC.html", gin.H{})
	// 		fmt.Println(c.PostForm("Distance0"))
	// 	})

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
