package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

func First(wg *sync.WaitGroup, streamSync [3]chan interface{}) {
	fmt.Printf("First")
	streamSync[0] <- nil
	wg.Done()
}

func Second(wg *sync.WaitGroup, streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Printf("Second")
	streamSync[1] <- nil
	wg.Done()
}

func Third(wg *sync.WaitGroup, streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Printf("Third")
	wg.Done()
}

func GetValueName(v interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
}

func PrintInOrde(callOrder [3]int) {
	// 取得輸入順序
	// var inputCallOrder []int
	// for _, arg := range os.Args[1:] {
	// 	num, _ := strconv.Atoi(arg)
	// 	inputCallOrder = append(inputCallOrder, num)
	// }
	inputCallOrder := callOrder
	fmt.Println("[]inputCallOrder:", inputCallOrder)

	// make an array of unbuffered
	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}

	// 建立 [int:func] 對應表
	var functionNumTable = map[int]func(*sync.WaitGroup, [3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}

	//依照輸入順序呼叫 goroutine
	wg := &sync.WaitGroup{}
	for _, fNum := range inputCallOrder {
		fmt.Println("Call:", GetValueName(functionNumTable[fNum]))
		wg.Add(1)
		go functionNumTable[fNum](wg, streamSync)
	}

	wg.Wait()
	fmt.Println()
}

func main() {
	var testCases = [][3]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, theCase := range testCases {
		startTime := time.Now()
		PrintInOrde(theCase)
		endTime := time.Now()
		fmt.Printf("Case %v Spent time:%v\n", theCase, endTime.Sub(startTime))
	}
}
