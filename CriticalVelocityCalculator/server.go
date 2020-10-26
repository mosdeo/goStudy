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

func CriticalVelocity(dinstance []int, raceTime []int)float32{
	Avg := func(nums []int)float32{
		var sum float32
		for _, num := range nums {
			sum += float32(num)
		}
		return sum/float32(len(nums))
	}

	//計算平均的最佳速度和距離
	avgDistance := Avg(dinstance)
	avgRaceTime := Avg(raceTime)

	//(T-MT)
	tAvgt := make([]float32, len(raceTime))
	for i, _ := range raceTime {
		tAvgt[i] = float32(raceTime[i]) - avgRaceTime
	} 
	
	//(D-MD)
	dAvgd := make([]float32, len(dinstance))
	for i, _ := range dinstance {
		tAvgt[i] = float32(dinstance[i]) - avgDistance
	}
	
	//Sum((T-MT)(D-MD))/Sum((T－MT)^2)
	//分子：numerator；分母：denominator
	var numeratorCV, denominatorCV, CV, minPerKmCV float32
	for i := range []int{0,1,2} {
		numeratorCV += dAvgd[i]*tAvgt[i]
		denominatorCV += tAvgt[i]*tAvgt[i]
	}
	CV = numeratorCV/denominatorCV
	minPerKmCV = (1000/CV)/60

	return minPerKmCV
}
