package main

import (
	"fmt"
	"math/rand"
	"time"
)

const mostDelayMicrosecond int = 99

func TrueRandom(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func P(num chan int, delayRate int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delayRate*TrueRandom(mostDelayMicrosecond)) * time.Microsecond)
		num <- i
		fmt.Println("生產了:", i)
	}
	fmt.Println("End of P()")
}

func C(num chan int, delayRate int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delayRate*TrueRandom(mostDelayMicrosecond)) * time.Microsecond)
		n, ok := <-num
		if ok {
			fmt.Println("消費了:", n)
		} else {
			fmt.Println("Channel is closed and empty.")
		}
	}
	fmt.Println("End of C()")
}

func main() {

	fmt.Println("=== Case1: Unbuffered ===")
	// Unbuffered 只能接受一吞接一吐，依序重複，正確步驟出現以前會先鎖死。
	nums := make(chan int)
	go P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)

	// Buffered 在順序上就很有彈性，比較快的生產者可以先把東西放入 Channel
	fmt.Println("=== Case2: Buffered ===")
	nums = make(chan int, 10)
	go P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)

	// Buffered 卻 Deadlock 的情況
	fmt.Println("=== Case3: Buffered but will Deadlock ===")
	nums = make(chan int, 9)
	P(nums, 1)
	go C(nums, 10)
	time.Sleep(2 * time.Second)
	// 原因是生產者 P() 沒有與消費者 C() 平行，所以 P() 把 channel 放滿以後就放不下去，
	// 此時 C() 尚未啟動，所以 channel 內的元素無法消化，於是同樣發生 Deadlock。
}
