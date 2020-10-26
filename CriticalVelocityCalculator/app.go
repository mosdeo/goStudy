package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// from uuid import uuid4
	uuid "github.com/satori/go.uuid"
)

func root(c *gin.Context) {
	uuidCookie, _ := c.Cookie("uuid")
	c.HTML(http.StatusOK, "index.html", gin.H{})
	c.SetCookie("uuid", uuidCookie, 3600, "/", "localhost", false, true)
}

func ask(c *gin.Context) {
	uuidCookie, err := c.Cookie("uuid")
	if err != nil {
		uuidCookie = "NotSet"
		u1 := uuid.Must(uuid.NewV4())
		c.SetCookie("uuid", u1.String(), 3600, "/", "localhost", false, true)
	}
	data := c.Request.PostFormValue("in")
	fmt.Printf("ip=%v, uuid=%v, data=%v\n", c.ClientIP(), uuidCookie, data)
	c.String(http.StatusOK, "Bot echo=[%v]", data)
}

func main() {
	router := gin.Default()

	router.StaticFile("/app.js", "./static/app.js")
	router.LoadHTMLFiles("static/index.html")

	router.GET("/", root)
	router.GET("/ask", ask)
	router.POST("/ask", ask)

	router.Run(":80")
}
