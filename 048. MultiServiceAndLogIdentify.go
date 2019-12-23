package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func Index01(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Welcome Server 01",
	})
}

func Index02(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Welcome Server 02",
	})
}

type LogFormatterWithInstanceIdentify struct {
	instanceName string
}

func (p *LogFormatterWithInstanceIdentify) defaultLogFormatterWithInstanceIdentify(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("[GIN] Service:%s |%v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
		p.instanceName,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

func main() {

	g.Go(func() error {
		router1 := gin.New()
		router1.Use(gin.Recovery())
		logFormatter := LogFormatterWithInstanceIdentify{instanceName: "01"}
		router1.Use(gin.LoggerWithFormatter(logFormatter.defaultLogFormatterWithInstanceIdentify))
		router1.GET("/", Index01)
		return router1.Run(":8080")
	})

	g.Go(func() error {
		router2 := gin.New()
		router2.Use(gin.Recovery())
		logFormatter := LogFormatterWithInstanceIdentify{instanceName: "02"}
		router2.Use(gin.LoggerWithFormatter(logFormatter.defaultLogFormatterWithInstanceIdentify))
		router2.GET("/", Index02)
		return router2.Run(":8081")
	})

	g.Wait()
}
