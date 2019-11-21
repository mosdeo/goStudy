package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	var First = func(wg *sync.WaitGroup) {
		fmt.Printf("First")
		wg.Done()
	}

	var Second = func(wg *sync.WaitGroup) {
		fmt.Printf("Second")
		wg.Done()
	}

	var Third = func(wg *sync.WaitGroup) {
		fmt.Printf("Third")
		wg.Done()
	}

	wg := &sync.WaitGroup{}
	var functionNumTable = map[int]func(*sync.WaitGroup){
		1: First,
		2: Second,
		3: Third,
	}

	// 取得輸入順序
	var inputCallOrder []int
	for _, arg := range os.Args[1:] {
		num, _ := strconv.Atoi(arg)
		inputCallOrder = append(inputCallOrder, num)
	}
	fmt.Println("inputCallOrder:", inputCallOrder)

	//依照輸入順序呼叫 goroutine

	for _, fNum := range inputCallOrder {
		wg.Add(1)
		go functionNumTable[fNum](wg)
	}

	wg.Wait()
	fmt.Println()
}
