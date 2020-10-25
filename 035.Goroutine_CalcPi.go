package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func PrintNumGoroutine() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

func main() {
	var termsNum float64 = 1E11
	var startTime, endTime time.Time

	startTime = time.Now()
	CalcPiByGoroutine(termsNum)
	endTime = time.Now()
	fmt.Println("CalcPiByGoroutine() Spent time:", endTime.Sub(startTime))

	startTime = time.Now()
	CalcPi(termsNum)
	endTime = time.Now()
	fmt.Println("CalcPi()            Spent time:", endTime.Sub(startTime))
}

func CalcPiByGoroutine(numOfTrem float64) {
	fmt.Println("runtime.NumCPU:", runtime.NumCPU())
	var SqrtPiDivieBy6 float64
	partSumStream := make(chan float64, runtime.NumCPU())
	wg := &sync.WaitGroup{}

	//每一個核心分配到的加總工作
	var GetPartSum = func(start float64, step float64, upper float64, sumStream chan<- float64, wg *sync.WaitGroup) {
		var thePartSum float64
		defer wg.Done() // return 前保證執行

		fmt.Println(start, step, upper)
		for i := start; i <= upper; i += step {
			thePartSum += 1.0 / (i * i)
			// fmt.Printf("start=%d, i=%f\n", start, i)
		}

		sumStream <- thePartSum
	}

	for coreNum := 1; coreNum <= runtime.NumCPU(); coreNum++ {
		wg.Add(1)
		go GetPartSum(float64(coreNum), float64(runtime.NumCPU()), numOfTrem-float64(runtime.NumCPU()-coreNum), partSumStream, wg)
	}

	PrintNumGoroutine()
	wg.Wait()
	close(partSumStream)
	SqrtPiDivieBy6 = func() (sum float64) {
		for part := range partSumStream {
			fmt.Println("part:", part)
			sum += part
		}
		return sum
	}()
	var pi = math.Pow(SqrtPiDivieBy6*6, 0.5)
	fmt.Printf("Pi=%f\n", pi)
}

func CalcPi(numOfTrem float64) {
	var SqrtPiDivieBy6 float64

	for i := 1.0; i <= numOfTrem; i++ {
		SqrtPiDivieBy6 += 1.0 / (i * i)
	}
	PrintNumGoroutine()

	var pi = math.Pow(SqrtPiDivieBy6*6, 0.5)
	fmt.Printf("Pi=%f\n", pi)
}
