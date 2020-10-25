package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const mostDelayMillisecond int = 99

func Print0(repeat int, wg *sync.WaitGroup) {
	defer wg.Done() //此 scope 結束前執行

	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		fmt.Printf("0")
	}
	fmt.Println("End of Print0()")
}

func Print1(repeat int, wg *sync.WaitGroup) {
	defer wg.Done() //此 scope 結束前執行

	for i := 0; i < repeat; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		fmt.Printf("1")
	}
	fmt.Println("End of Print1()")
}

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func main() {
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(2)

	//wg.Wait() // Before goroutine: fatal error: all goroutines are asleep - deadlock!
	go Print0(99, wg)
	go Print1(99, wg)
	wg.Wait() // Have to after goroutine
}
