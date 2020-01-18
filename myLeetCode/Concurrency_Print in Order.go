package main

import (
	"fmt"
)

func First(streamSync [3]chan interface{}) {
	fmt.Print("First")
	streamSync[0] <- nil
}

func Second(streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Print("Second")
	streamSync[1] <- nil
}

func Third(streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Print("Third")
	streamSync[2] <- nil
}

func PrintInOrde(callOrder [3]int) {
	inputCallOrder := callOrder
	fmt.Println("[]inputCallOrder:", inputCallOrder)

	// make an array of unbuffered
	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}

	// 建立 [int:func] 對應表
	var functionNumTable = map[int]func([3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}

	//依照輸入順序呼叫 goroutine
	for _, fNum := range inputCallOrder {
		go functionNumTable[fNum](streamSync)
	}

	<-streamSync[2]
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
		PrintInOrde(theCase)
		fmt.Println()
		fmt.Println()
	}
}
