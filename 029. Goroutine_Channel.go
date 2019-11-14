package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mostDelayMillisecond int = 1

func P(num chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		num <- i
		fmt.Println("生產了:", i)
	}
	fmt.Println("End of P()")
}

func C(num chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(TrueRandom(mostDelayMillisecond)) * time.Millisecond)
		n, ok := <-num
		if ok {
			fmt.Println("消費了:", n)
		} else {
			fmt.Println("Channel is closed and empty.")
		}
	}
	fmt.Println("End of C()")
}

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func main() {
	nums := make(chan int)
	go P(nums)
	go C(nums)
	time.Sleep(10 * time.Second)
}
